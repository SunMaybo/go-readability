package readability

import (
	"fmt"
	"io/ioutil"
	"os"
	fp "path/filepath"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
	"golang.org/x/net/html"
	"bytes"
	"net/http"
	"log"
)


func TestSource(t *testing.T){
	var textList []string
	textList = append(textList, `<span>来源：www.bstsjx.com</span> `)
	textList = append(textList, `<div class="info">
              <span>作者：股指模拟开户</a></span><span>   来源：www.hesyy666.com</span><span>   时间：2019年06月03日 10:32</span><span>   点击：<script src="/plus/count.php?view=yes&amp;aid=3161&amp;mid=1" type="text/javascript" language="javascript"></script>370</span><span style="display:none;"></span>
            </div>`)
	textList = append(textList, `
                <span class="source">本文来源于:每日经济新闻</span></span>
<p>图片来源：视觉中国</p>
`)
	textList = append(textList, `
                <span class="source">本文来自:每日经济新闻</span></span>
<p>图片来源：视觉中国</p>
`)
	textList = append(textList, `	并 本报北京6月5日电 中国青年报·中青在线记者樊未晨来源：中国青年报。
	先跟考生所在中学或各省级招生考试机构公</p>
`)
	for _, text := range textList {
		textList = append(textList, `从欧盟出台《通用数据保护条例》（GDPR）实施“严监管”，到
图片来源：图虫创意
据参考消息报道，20国集团（G20）财长会议6月8日在日本福冈举行，会议支持对亚马逊、谷歌和脸书等数字巨头征税，使得这个建议在全球范围内更加接近于可行。`)
		p := Parser{}
		source:=p.FilterSourceName(text)
		fmt.Println(source)
	}
}
var content = `

<!DOCTYPE html>
<html>

<head>
    <title>被红岭起诉 长城资产回应：重组方已进驻会先付-投资学院-送钱网</title>
    <meta name="keywords" content="被红岭起诉 长城资产回应：重组方已进驻会先付-投资学院-送钱网" />
    <meta name="description" content="红岭创投董事长周世平有点怒了。-投资学院-送钱网">
    	<meta name="author" content="YinQiao">
	<meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    <!-- Set render engine for 360 browser -->
    <meta name="renderer" content="webkit">
   	<!-- No Baidu Siteapp-->
    <meta http-equiv="Cache-Control" content="no-siteapp"/>
    <!-- HTML5 shim for IE8 support of HTML5 elements -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
    <![endif]-->
	<link rel="icon" href="/themes/simplebootx/Public/images/favicon.ico" type="image/x-icon">
    <link rel="shortcut icon" href="/themes/simplebootx/Public/images/favicon.ico" type="image/x-icon">
    <link rel="stylesheet" href="/themes/simplebootx/Public/css/common.css" />
    <link rel="stylesheet" href="/themes/simplebootx/Public/css/college.css" />
    <link rel="stylesheet" href="/themes/simplebootx/Public/css/index.css">
    <link rel="stylesheet" href="/themes/simplebootx/Public/css/list.css">   
    <link rel="stylesheet" href="/themes/simplebootx/Public/css/common.83286230.css"/>
   <link rel="stylesheet" href="/themes/simplebootx/Public/css/home.1d2f145b.css"/>
    
    <script type="text/javascript">
        var App= "";
    </script>
	<!--[if IE 7]>
	<link rel="stylesheet" href="/themes/simplebootx/Public/simpleboot/font-awesome/4.4.0/css/font-awesome-ie7.min.css">
	<![endif]-->
    <script src="/themes/simplebootx/Public/js/plugins/jquery-1.12.1.js"></script>
    <!-- <script type="text/javascript" src="/themes/simplebootx/Public/js/jquery.js"></script> -->
     <script type="text/javascript" src="/themes/simplebootx/Public/js/layer.js"></script>	
	 <script type="text/javascript">
	         var controllername="College";
	         var actionname="detail";
	         var getid="";
	         function browserRedirect() {
	             var sUserAgent = navigator.userAgent.toLowerCase();
	             var bIsIpad = sUserAgent.match(/ipad/i) == "ipad";
	             var bIsIphoneOs = sUserAgent.match(/iphone os/i) == "iphone os";
	             var bIsMidp = sUserAgent.match(/midp/i) == "midp";
	             var bIsUc7 = sUserAgent.match(/rv:1.2.3.4/i) == "rv:1.2.3.4";
	             var bIsUc = sUserAgent.match(/ucweb/i) == "ucweb";
	             var bIsAndroid = sUserAgent.match(/android/i) == "android";
	             var bIsCE = sUserAgent.match(/windows ce/i) == "windows ce";
	             var bIsWM = sUserAgent.match(/windows mobile/i) == "windows mobile";
	             if (bIsIpad || bIsIphoneOs || bIsMidp || bIsUc7 || bIsUc || bIsAndroid || bIsCE || bIsWM) {
	                 if(getid){
	                     var url="/m/"+controllername+"/"+getid+".html";
	                 }else{
	                     var url="/m/"+controllername+"/"+actionname+".html";
	                 }
	                 window.location.href = url;
	             }
	         }
	         browserRedirect();
	     </script>
</head>

<body>
         
        <link href="/themes/simplebootx/Public/css/style.css" rel="stylesheet">
        
	<script type="text/javascript" src="/themes/simplebootx/Public/js/bootstrap.js"></script>
        <script type="text/javascript" src="/themes/simplebootx/Public/js/style.js"></script>
<!--头部通用-->
<header>
	<!--右侧浮动-->
	<div class="float">
		<div class="erweima" style="">
		</div>
		<div class="kefu" style="display: none;">
			<dl>
				<dt class="kefu_title">在线客服</dt>
				
				<!-- <dd><a class="kefu_dd" target="_blank" href="http://wpa.qq.com/msgrd?v=3&amp;uin=79189952&amp;site=qq&amp;menu=yes">客服一号</a></dd> -->
			
				<dd><a class="kefu_dd" target="_blank" href="http://wpa.qq.com/msgrd?v=3&amp;uin=531247347&amp;site=qq&amp;menu=yes">客服一号</a></dd>
			
				<dd><a class="kefu_dd" target="_blank" href="http://wpa.qq.com/msgrd?v=3&amp;uin=531247347&amp;site=qq&amp;menu=yes">客服二号</a></dd>
			
				<dd><a class="kefu_dd" target="_blank" href="http://wpa.qq.com/msgrd?v=3&amp;uin=1828692306&amp;site=qq&amp;menu=yes">商务合作</a></dd>
			</dl>
		</div>
		<ul id="toTopUl">
			<li><a class="flo_1" target="_blank" href="/aboutus/scafina.html" style="font-size: 12px;"><img src="/themes/simplebootx/Public/images/home/icon/float_5.png">收益计算器</a></li>
			<li><a class="flo_2" href="javascript:void(0);" style="font-size: 12px;"><img src="/themes/simplebootx/Public/images/home/icon/float_7.png" style="margin-right: 0px; ">0530-5342994</a></li>
			<li><a class="flo_3" href="javascript:void(0);"><img src="/themes/simplebootx/Public/images/home/icon/float_6.png"></a></li>
			<li><a class="flo_w" href="javascript:void(0);" onmouseover="showCode();" onmouseout="hideCode();"><img src="/themes/simplebootx/Public/images/home/icon/float_9.png"></a></li>
		</ul>
	</div>
	<!--右侧浮动结束-->
	<div class="header_top">
		<div class="header_center clearfix">
			<div class="header_lef">
				<span class="mr15"><i class="glyphicon glyphicon-earphone"></i>欢迎致电：0530-5342994</span>
				<span><i class="glyphicon glyphicon-time"></i>服务时间：9:00-21:00</span>
				<span id="qq">
					<i class="icon icon_qq"></i><i>送钱网官方交流群II群</i>
					<em class="show_qq">
						<p>送钱网官方交流群</p>
						<p><a href="//shang.qq.com/wpa/qunwpa?idkey=7456307f5236a04aaad6b05abe978d1dd1cef4b6b29e4f5bd9f3819021267fbd" target="_blank">QQ群：669600739</a></p>
					</em>
				</span>	
			</div>
			<div class="header_rig">
				<ul>
					<li><a href="/notice.html" target="_blank" title="官方公告">官方公告</a></li>
					<li><a href="/aboutus/help.html" target="_blank">帮助中心</a></li>
					<li><a href="/aboutus/index.html" target="_blank">关于我们</a></li>
                                         
                                            <li><a href="/user/login.html" target="_blank">登录</a></li>
                                            <li><a href="/user/register.html" target="_blank">注册</a></li>                                </ul>
			</div>
		</div>
	</div>
	<div class="header_bottom">
		<div class="bottom_center">
			<a href="/"><img class="bot_left" src="/themes/simplebootx/Public/images/LOGO.png" alt="送钱网" title="送钱网"></a>	
			<div class="nav_list">
				<ul>
					<li><a href="/">首页</a></li>
					<li><a href="/witkey.html">首投返利</a>
					</li>
					<li><a href="/rebate.html">复投返利</a></li>
					<li><a href="/college.html">理财学院</a></li>
					<li><a href="/activity.html">活动专区</a><i class="hot"><img src="/themes/simplebootx/Public/images/bgicon.gif"></i></li>
                                                                                 <li><a href="/user/login.html"><i class="account mr15"><img src="/themes/simplebootx/Public/images/home/icon/user_head.png"></i>我的账户</a>
                                        </li>				</ul>
			</div>
			
		</div>
	</div>

</header>
<script type="text/javascript">
	$(function(){
		$('#toTopUl > li').hover(function(){
			$(this).find('.flo_1 > img').attr('src','/themes/simplebootx/Public/images/home/icon/float_55.png');
			$(this).find('.flo_1').stop().animate({"width":"170px"},200).css({"opacity":"1","filter":"Alpha(opacity=100)","background":"#67bad4"});
			$(this).find('.flo_2').stop().animate({"width":"170px"},200).css({"opacity":"1","filter":"Alpha(opacity=100)","background":"#67bad4"});
		},function(){
			$(this).find('.flo_1 > img').attr('src','/themes/simplebootx/Public/images/home/icon/float_5.png');
			$(this).find('.flo_1').stop().animate({"width":"52px"},200).css({"opacity":"1","filter":"Alpha(opacity=100)","background":"none"});
			$(this).find('.flo_2').stop().animate({"width":"52px"},200).css({"opacity":"1","filter":"Alpha(opacity=100)","background":"none"});
		});
	});
</script>
<!--头部结束-->
<!-- /头部 -->
    <!---->
    <div class="college-container">
        <div class="wrapper">
            <div class="crumbs">
                您的位置:
                <a href="/" target="_blank">首页</a> >
                <a href="/college.html">投资学院</a> >
                <a href="/college/classify/7.html">金融</a> >
                <a href="javascript:;" class="active">被红岭起诉 长城资产回应：重组方已进驻会先付</a>
            </div>
            <div class="college-main clearfix">
                <div class="college-left">
                    <div class="college-detail">
                        <div class="news_bt">
                            <h3>被红岭起诉 长城资产回应：重组方已进驻会先付</h3></div>
                        <div class="show_tag">
                            <span>2019-06-13 16:01:38</span>
                            <div class="share">
                                <div class="bdsharebuttonbox">
                                    <a href="#" class="bds_more" data-cmd="more"></a>
                                    <a href="#" class="bds_qzone" data-cmd="qzone" title="分享到QQ空间"></a>
                                    <a href="#" class="bds_tsina" data-cmd="tsina" title="分享到新浪微博"></a>
                                    <a href="#" class="bds_tqq" data-cmd="tqq" title="分享到腾讯微博"></a>
                                    <a href="#" class="bds_renren" data-cmd="renren" title="分享到人人网"></a>
                                    <a href="#" class="bds_weixin" data-cmd="weixin" title="分享到微信"></a>
                                </div>
                                <script>
                                window._bd_share_config = {
                                    "common": {
                                        "bdSnsKey": {},
                                        "bdText": "",
                                        "bdMini": "2",
                                        "bdMiniList": false,
                                        "bdPic": "",
                                        "bdStyle": "1",
                                        "bdSize": "24"
                                    },
                                    "share": {},
                                    "image": {
                                        "viewList": ["qzone", "tsina", "tqq", "renren", "weixin"],
                                        "viewText": "分享到：",
                                        "viewSize": "16"
                                    },
                                    "selectShare": {
                                        "bdContainerClass": null,
                                        "bdSelectMiniList": ["qzone", "tsina", "tqq", "renren", "weixin"]
                                    }
                                };
                                with(document) 0[(getElementsByTagName('head')[0] || body).appendChild(createElement('script')).src = 'http://bdimg.share.baidu.com/static/api/js/share.js?v=89860593.js?cdnversion=' + ~(-new Date() / 36e5)];
                                </script>
                                <!--   <a href="#" title=" 微信" class="weixin"></a>
                                <a href="#" title="微博" class="weibo"></a>
                                <a href="#" title="qq空间" class="qq"></a> -->
                            </div>
                        </div>
                        <div class="detail-main">
                            红岭创投董事长周世平有点怒了。                                                        <div style="padding:10px;text-align: center">
                                    <img src="/data/upload/20190613/5d0202dfec99b.jpg" width="500" height="310" alt="被红岭起诉 长城资产回应：重组方已进驻会先付" />
                                </div>                            <p>继公开“催债”之后，周世平开怼四大AMC之一的某资产管理公司，称即日启动诉讼。</p><p><br/></p><p>6月12日，周世平对外称，根据某资产管理公司与红岭创投的合约，2018年6月28日正式到期，本金三亿元及利息、罚息等至今未能处理，红岭创投本着友好协商的态度多次沟通，但至今未有明显进展。“贵为四大之一，这不是你应该有的态度，即日起启动诉讼程序，対不起，曾经耐心等过你。”</p><p><br/></p><p>虽然周世平未具体点名这个四大之一的资管公司，这个被催债的资产管理公司正是长城资产内蒙古分公司。</p><p><br/></p><p>长城资产分公司回应：已正式由第三方重组接盘</p><p><br/></p><p>早在4月9日，周世平在官网发布“催债”帖称，2017年6月，在内蒙古呼和浩特市当地政府的推动下，开发商内蒙古联发房地产在红岭创投借款3亿元，综合利率为年化15％，借款期限12个月，到期一次性归还剩余本金，并引入了担保方中国长城资产内蒙古分公司为此项目兜底回购。</p><p><br/></p><p>但是，内蒙古联发房地产开发有限公司在红岭创投的贷款已经出现逾期。截至2019年3月25日，贷款本金余额3亿元，利息加罚息共计652.87万元，利息、罚息和滞纳金共计5197.04万元。周世平此前已表示红岭创投在正准备起诉中国长城资产内蒙古分公司，确保债权收回。</p><p><br/></p><p>中国长城资产内蒙古分公司此前回应称，“项目逾期后，由于房地产监管政策的变化，公司收购这个房地产项目的方式已经难以实现。</p><p><br/></p><p>从今年4月到现在，可能项目处理没有进展，周世平才对长城资产发起诉讼。</p><p><br/></p><p>6月13日，长城资产内蒙古分公司有关负责人紧急回应《华夏时报》记者采访时说，“我们和红岭创投老周保持着沟通。这个房地产项目，现已正式由第三方重组接盘，目前第三方已进驻，会优先支付红岭创投5000万元，等项目全部完工之后，还有十几亿的销售回款，也会优先支付给红岭创投。如果回收尾款不理想，第三方也会支付红岭创投剩下的钱。”</p><p><br/></p><p>该负责人解释说，其实欠钱的不是我们，而是内蒙联发房地产，他们融资，我们做回购，收购这笔债权，因为现在政策有调整，收不上来，但这个过程中，我们积极努力，帮助房地产企业引进第三方重组，现在已签订了托管协议。</p><p><br/></p><p>一名金融机构人士表示，按常规来说，既然已签订协议担保回购，不管什么政策原因还是市场等原因，都应该履行合同，由担保方兜底负责。</p><p><br/></p><p>作为行业最大的P2P平台之一，红岭创投宣布“三年退出计划”之后，周世平一直在为“钱”奔波，甚至直言平台现在急需解决的核心就是“收债、收债、收债”，以满足流动性需求。</p><p><br/></p><p>5月29日，周世平发帖称，平台线上线下总兑付本金规模260亿，目前重组洽谈进展顺利，有望在上半年正式形成落地方案。他同时直言，现在红岭创投的兑付方案已经确定，根据可行性,今年能够收回的资金，在40亿元到50亿元之间，这里面包含正常的项目回款和不良项目的清收等。</p><p><br/></p><p>“红岭系”各个网贷，包括红岭创投、亿钱贷等平台线上线下总兑付本金规模260亿，这个待偿金额确实不小。截至目前，红岭创投共对投资者进行了八次兑付，合计兑付了6亿。</p><p><br/></p><p>6月2日，周世平发贴回应投资者的诸多问题中，红岭系各平台重组的进展仍然备受关注。周世平表示，虽然面临各种困难，而且今年二、三季度应该是最困难的时期，各种因素内外交困，但是不应该只顾眼前的困难去贱卖资产。选择三年清盘，引进机构重组更多是为今后考虑，将资产有序清理变现，发挥其最大的价值才能让清盘更具安全性、确定性。</p><p><br/></p><p>股票平仓危机</p><p><br/></p><p>屋漏偏逢连夜雨。</p><p><br/></p><p>在200多亿元的待偿压力下，一边各类资产遭遇回收难，另一边则是迫在眉睫的强制平仓。</p><p><br/></p><p>6月13日，深南股份（002417.SZ）发布公告称，公司大股东周世平于2019年6月11日继续被动减持3042000股，占公司总股本的1.13%，成交金额19991490元。周世平质押的部分公司股份存在被质权人首创证券进行处置而导致继续被动减持的可能。</p><p><br/></p><p>公告称，因股权质押业务被进行处置而导致被动减持的股份数量合计为6016900股，占公司总股本的2.23%，成交金额合计为人民币39968344.45元。本次减持后，周世平先生及其一致行动人红岭控股有限公司合计持有公司股份73663214股，占公司总股本的27.28%，周世平先生仍为公司第一大股东及实际控制人。</p><p><br/></p><p>此前6月3日，深南股份披露大股东周世平所持股份可能遭到首创证券强行平仓的公告，股价遭遇了开盘跌停。这份公告显示，周世平所持该公司股份因股票质押回购业务，近期可能将遭首创证券被动减持12.81%。</p><p><br/></p><p>据了解，周世平是在2016年7月在首创证券办理了股票质押式回购交易业务，质押到期日为2019年7月12日。6月2日，周世平在红岭创投社区公开回复出借人称，周世平质押股份总数为3458万股，融资金额为1亿元人民币，质押式回购合约“并未触及平仓线”；深南股份公司主营业务经营正常，并购正常推进中，即将迎来业绩拐点；红岭系各平台已初步稳定而且跟上市公司有防火墙隔离，不会传导风险。</p><p><br/></p><p>深南股份此前表示，周世平仍在积极与首创证券进行沟通，努力通过在合约到期前购回等相关措施，尽力避免或降低本次可能被动减持及可能发生违规减持的不利影响，妥善解决可能被动减持的问题。</p><p><br/></p><p>目前，深南股份的股价一直处于低位徘徊。在2016年7月，周世平签订股份质押回购合约时，深南股份股价仍在15元左右，三年后，深南股份的股价已经跌至6.5元，跌去将近6成，且一直处在“保壳边缘”。</p><p><br/></p><p>2018年年报显示，深南股份营收为1.81亿元，较上年同期的1.05亿元增长73.30%，但是归属股东的净利润为-4906.78万元，较上年同期的767.18万元下降739.58%。</p><p><br/></p><p>目前，上市公司的平仓压力，还有P2P大量不良资产待处置，“老周”目前身上的还钱压力可谓不小。一名市场人士表示，红岭创投是曾经最具影响力的互金公司，老周也是最阳光的老板，红岭创投解决眼下困难，走出危机对稳定市场都有重要意义。</p><p><br/></p>                        </div>
                        <div class="statement">【声明】本文内容由网友发布，仅代表网友个人经验或观点，不代表本网站立场和观点。如果本文侵犯了您的知识产权，请 与我们取得联系，我们会及时修改或删除
                        </div>
                        <div class="source">
                            <p>（文章来源：<a href="javascript:void(0);"><a href="http://hjcyh.com/college/7876.html">http://hjcyh.com/college/7876.html</a></a>）</p>
                        </div>
                        <div class="paging">
                            <p>上一篇：
                                <a href="/college/7877.html">胡玮炜挥别摩拜单车 正式卸任法定代表人</a>
                                                            </p>
                            <p>下一篇：
                                <a href="/college/7875.html">伽满优宣布良性清盘，兑付方案分24期</a>
                                                            </p>
                        </div>
                    </div>
                    <!--列表-->
                    <div class="p2p-list">
                        <div class="p2pBox">
                            <div class="hd">
                                <ul>
                                    <li class="on">P2P</li><li class="on">平台</li><li class="on">数据</li><li class="on">政策</li><li class="on">金融</li><li class="on">理财</li><li class="on">融资</li><li class="on">利率</li><li class="on">曝光</li>                                </ul>
                            </div>
                            <div class="bd">
                                <ul>
                                    <li style="border:0px;padding:0px;">
                                        <div class="post-left">
                                            <div class="left-img">
                                                <a href="/college/8076.html" target="_blank" title="宜贷网发布致歉信:兑付不尽人意 需要更多时间">
                                                                                                    <img src="/data/upload/20190708/5d22f37d88cb7.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>宜贷网发布致歉信:兑付不尽人意 需要更多时间</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div><div class="left-img">
                                                <a href="/college/8073.html" target="_blank" title="网信发布最新通知：延迟原定下周一的投资人见面会">
                                                                                                    <img src="/data/upload/20190708/5d2292a159458.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>网信发布最新通知：延迟原定下周一的投资人见面会</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div>                                        </div>
                                        <div class="post-right">
                                            <ul>
                                                <li><a href="/college/8076.html" target="_blank" title="宜贷网发布致歉信:兑付不尽人意 需要更多时间">宜贷网发布致歉信:兑付不尽人意 需要更多时间</a></li><li><a href="/college/8073.html" target="_blank" title="网信发布最新通知：延迟原定下周一的投资人见面会">网信发布最新通知：延迟原定下周一的投资人见面会</a></li><li><a href="/college/8069.html" target="_blank" title="宜信公告：网传图片系“造谣” 已找到造谣者信息并报警">宜信公告：网传图片系“造谣” 已找到造谣者信息并报警</a></li><li><a href="/college/8066.html" target="_blank" title="众车在线最新进展 新增1500万到账">众车在线最新进展 新增1500万到账</a></li><li><a href="/college/8065.html" target="_blank" title="网信再发公告：正常运营 集团母公司在积极帮助兑付">网信再发公告：正常运营 集团母公司在积极帮助兑付</a></li><li><a href="/college/8062.html" target="_blank" title="宜贷网：兑付超7.5亿 回款占待偿本金的25%">宜贷网：兑付超7.5亿 回款占待偿本金的25%</a></li><li><a href="/college/8057.html" target="_blank" title="30家立案平台最新案情进展梳理">30家立案平台最新案情进展梳理</a></li><li><a href="/college/8054.html" target="_blank" title="94家平台4月协会信披盘点：拍拍贷未及时更新数据">94家平台4月协会信披盘点：拍拍贷未及时更新数据</a></li><li><a href="/college/8053.html" target="_blank" title="网信理财 还是扛不住了！">网信理财 还是扛不住了！</a></li><li><a href="/college/8036.html" target="_blank" title="网贷平台备案试点继续延期 6月成交量同比锐减逾5成">网贷平台备案试点继续延期 6月成交量同比锐减逾5成</a></li>                                            </ul>
                                        </div>
                                    </li>
                                </ul><ul>
                                    <li style="border:0px;padding:0px;">
                                        <div class="post-left">
                                            <div class="left-img">
                                                <a href="/college/8028.html" target="_blank" title="警方通报钱眼金融案情进展：查封24套房产">
                                                                                                    <img src="/data/upload/20190702/5d1ab00d684a5.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>警方通报钱眼金融案情进展：查封24套房产</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div><div class="left-img">
                                                <a href="/college/8026.html" target="_blank" title="天玑财富被警方立案侦查 3人被刑事拘留">
                                                                                                    <img src="/data/upload/20190702/5d1aacd4b11f5.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>天玑财富被警方立案侦查 3人被刑事拘留</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div>                                        </div>
                                        <div class="post-right">
                                            <ul>
                                                <li><a href="/college/8028.html" target="_blank" title="警方通报钱眼金融案情进展：查封24套房产">警方通报钱眼金融案情进展：查封24套房产</a></li><li><a href="/college/8026.html" target="_blank" title="天玑财富被警方立案侦查 3人被刑事拘留">天玑财富被警方立案侦查 3人被刑事拘留</a></li><li><a href="/college/7946.html" target="_blank" title="P2P大时贷案2名嫌犯已移送审查起诉">P2P大时贷案2名嫌犯已移送审查起诉</a></li><li><a href="/college/7810.html" target="_blank" title="微贷网2019年一季度营收9.5亿 净利润1.35亿">微贷网2019年一季度营收9.5亿 净利润1.35亿</a></li><li><a href="/college/7718.html" target="_blank" title="今金贷案进展：张彦军等6人被捕 查封房产30套">今金贷案进展：张彦军等6人被捕 查封房产30套</a></li><li><a href="/college/7631.html" target="_blank" title="红岭创投发布资产收益权转让意见征集稿">红岭创投发布资产收益权转让意见征集稿</a></li><li><a href="/college/7562.html" target="_blank" title="抱财网旗下资产端年亏5000万业绩对赌失败">抱财网旗下资产端年亏5000万业绩对赌失败</a></li><li><a href="/college/7538.html" target="_blank" title="揭24家P2P大平台最新数据：8家近一年逾期率走势曝光(名单)">揭24家P2P大平台最新数据：8家近一年逾期率走势曝光(名单)</a></li><li><a href="/college/7495.html" target="_blank" title="深圳P2P绿化贷案有新进展 如期还款7962余万">深圳P2P绿化贷案有新进展 如期还款7962余万</a></li><li><a href="/college/7461.html" target="_blank" title="有融网案情进展：警方已开通借款人还款账户">有融网案情进展：警方已开通借款人还款账户</a></li>                                            </ul>
                                        </div>
                                    </li>
                                </ul><ul>
                                    <li style="border:0px;padding:0px;">
                                        <div class="post-left">
                                            <div class="left-img">
                                                <a href="/college/6530.html" target="_blank" title="“2018全球金融科技100强”榜单发布：蚂蚁金服第一，京东金融第二">
                                                                                                    <img src="/data/upload/20181026/5bd2bf29d43b7.png" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>“2018全球金融科技100强”榜单发布：蚂蚁金服第一，京东金融第二</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div><div class="left-img">
                                                <a href="/college/6247.html" target="_blank" title="P2P恶意逃废债企业和个人信息将纳入征信系统 百万条信息已上报">
                                                                                                    <img src="/data/upload/20180903/5b8ce370e0cf7.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>P2P恶意逃废债企业和个人信息将纳入征信系统 百万条信息已上报</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div>                                        </div>
                                        <div class="post-right">
                                            <ul>
                                                <li><a href="/college/6530.html" target="_blank" title="“2018全球金融科技100强”榜单发布：蚂蚁金服第一，京东金融第二">“2018全球金融科技100强”榜单发布：蚂蚁金服第一，京东金融第二</a></li><li><a href="/college/6247.html" target="_blank" title="P2P恶意逃废债企业和个人信息将纳入征信系统 百万条信息已上报">P2P恶意逃废债企业和个人信息将纳入征信系统 百万条信息已上报</a></li><li><a href="/college/6246.html" target="_blank" title="P2P平台逾期风暴持续之下，上市公司各显神通为逾期的项目埋单">P2P平台逾期风暴持续之下，上市公司各显神通为逾期的项目埋单</a></li><li><a href="/college/6245.html" target="_blank" title="8月网贷正常运营平台持续下降1595家，刷新近一年以来的新高">8月网贷正常运营平台持续下降1595家，刷新近一年以来的新高</a></li><li><a href="/college/6230.html" target="_blank" title="余额宝的用户数已经接近美国总人口的2倍 ">余额宝的用户数已经接近美国总人口的2倍 </a></li><li><a href="/college/6225.html" target="_blank" title="真格系火币、币世界“踩雷”">真格系火币、币世界“踩雷”</a></li><li><a href="/college/6147.html" target="_blank" title="央行报纸：禁止地方政府变相鼓励逃废银行债务等行为">央行报纸：禁止地方政府变相鼓励逃废银行债务等行为</a></li><li><a href="/college/6105.html" target="_blank" title="乱象犹存 金交所再遭摸底清查">乱象犹存 金交所再遭摸底清查</a></li><li><a href="/college/6084.html" target="_blank" title="抓老赖现场曝光 100多名干警出动">抓老赖现场曝光 100多名干警出动</a></li><li><a href="/college/6077.html" target="_blank" title="深圳互金协会公布首批9名P2P老赖名单 将上报并纳入征信">深圳互金协会公布首批9名P2P老赖名单 将上报并纳入征信</a></li>                                            </ul>
                                        </div>
                                    </li>
                                </ul><ul>
                                    <li style="border:0px;padding:0px;">
                                        <div class="post-left">
                                            <div class="left-img">
                                                <a href="/college/7554.html" target="_blank" title="北京广东山东取消企业银行账户许可 开户手续更便捷">
                                                                                                    <img src="/data/upload/20190429/5cc657566cf9b.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>北京广东山东取消企业银行账户许可 开户手续更便捷</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div><div class="left-img">
                                                <a href="/college/7505.html" target="_blank" title="央行：新版征信采集水电缴费征信数据需获授权">
                                                                                                    <img src="/data/upload/20190423/5cbe6ac7215e3.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>央行：新版征信采集水电缴费征信数据需获授权</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div>                                        </div>
                                        <div class="post-right">
                                            <ul>
                                                <li><a href="/college/7554.html" target="_blank" title="北京广东山东取消企业银行账户许可 开户手续更便捷">北京广东山东取消企业银行账户许可 开户手续更便捷</a></li><li><a href="/college/7505.html" target="_blank" title="央行：新版征信采集水电缴费征信数据需获授权">央行：新版征信采集水电缴费征信数据需获授权</a></li><li><a href="/college/7389.html" target="_blank" title="税务总局5月将再推服务措施降低社保费率">税务总局5月将再推服务措施降低社保费率</a></li><li><a href="/college/7294.html" target="_blank" title="《商业银行资产托管业务指引》">《商业银行资产托管业务指引》</a></li><li><a href="/college/7194.html" target="_blank" title="厦门市金融工作办公室关于进一步开展融资租赁公司摸底排查的通知">厦门市金融工作办公室关于进一步开展融资租赁公司摸底排查的通知</a></li><li><a href="/college/7053.html" target="_blank" title="《关于办理非法集资刑事案件若干问题的意见》">《关于办理非法集资刑事案件若干问题的意见》</a></li><li><a href="/college/6968.html" target="_blank" title="《东莞市商品房认购书（范本）》的通知》">《东莞市商品房认购书（范本）》的通知》</a></li><li><a href="/college/6899.html" target="_blank" title="年终奖个税新规 算算省了多少钱">年终奖个税新规 算算省了多少钱</a></li><li><a href="/college/6848.html" target="_blank" title="《关于加强存款准备金管理有关事项的通知》">《关于加强存款准备金管理有关事项的通知》</a></li><li><a href="/college/6796.html" target="_blank" title="深圳发布小额贷款保证保险试点实施办法">深圳发布小额贷款保证保险试点实施办法</a></li>                                            </ul>
                                        </div>
                                    </li>
                                </ul><ul>
                                    <li style="border:0px;padding:0px;">
                                        <div class="post-left">
                                            <div class="left-img">
                                                <a href="/college/8081.html" target="_blank" title="3家上市公司将集体退市，竟有人还在抄底">
                                                                                                    <img src="/data/upload/20190708/5d2305376e7d0.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>3家上市公司将集体退市，竟有人还在抄底</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div><div class="left-img">
                                                <a href="/college/8080.html" target="_blank" title="招行回应异常利率交易事件：已暂停了相关交易员资格">
                                                                                                    <img src="/data/upload/20190708/5d2302d7818d1.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>招行回应异常利率交易事件：已暂停了相关交易员资格</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div>                                        </div>
                                        <div class="post-right">
                                            <ul>
                                                <li><a href="/college/8081.html" target="_blank" title="3家上市公司将集体退市，竟有人还在抄底">3家上市公司将集体退市，竟有人还在抄底</a></li><li><a href="/college/8080.html" target="_blank" title="招行回应异常利率交易事件：已暂停了相关交易员资格">招行回应异常利率交易事件：已暂停了相关交易员资格</a></li><li><a href="/college/8079.html" target="_blank" title="平安银行回应异常利率交易事件：已暂停相关交易员权限">平安银行回应异常利率交易事件：已暂停相关交易员权限</a></li><li><a href="/college/8077.html" target="_blank" title="拿去花资产暴涨背后保理ABS不良大增 携程能否复制花呗小贷？">拿去花资产暴涨背后保理ABS不良大增 携程能否复制花呗小贷？</a></li><li><a href="/college/8075.html" target="_blank" title="排放造假被罚1.7亿元 江淮汽车大跌逾6.07％">排放造假被罚1.7亿元 江淮汽车大跌逾6.07％</a></li><li><a href="/college/8074.html" target="_blank" title=" 德银耗资74亿欧元整改：退出股票交易业务 拟裁员2万人"> 德银耗资74亿欧元整改：退出股票交易业务 拟裁员2万人</a></li><li><a href="/college/8072.html" target="_blank" title="二三四五设立互联网保险公司未果 与各方协商后退出">二三四五设立互联网保险公司未果 与各方协商后退出</a></li><li><a href="/college/8071.html" target="_blank" title="51公积金IPO搁浅">51公积金IPO搁浅</a></li><li><a href="/college/8070.html" target="_blank" title="ST慧球称中江信托《保兑函》所盖上市公司公章系伪造">ST慧球称中江信托《保兑函》所盖上市公司公章系伪造</a></li><li><a href="/college/8068.html" target="_blank" title="新城控股被多家评级机构列入负面观察名单">新城控股被多家评级机构列入负面观察名单</a></li>                                            </ul>
                                        </div>
                                    </li>
                                </ul><ul>
                                    <li style="border:0px;padding:0px;">
                                        <div class="post-left">
                                            <div class="left-img">
                                                <a href="/college/8078.html" target="_blank" title="交行发售理财子公司产品：1元起投 股市投资比例最高20％">
                                                                                                    <img src="/data/upload/20190708/5d22fe043d968.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>交行发售理财子公司产品：1元起投 股市投资比例最高20％</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div><div class="left-img">
                                                <a href="/college/8014.html" target="_blank" title="联创永宣陷股东退出纠纷">
                                                                                                    <img src="/data/upload/20190629/5d1709fd26570.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>联创永宣陷股东退出纠纷</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div>                                        </div>
                                        <div class="post-right">
                                            <ul>
                                                <li><a href="/college/8078.html" target="_blank" title="交行发售理财子公司产品：1元起投 股市投资比例最高20％">交行发售理财子公司产品：1元起投 股市投资比例最高20％</a></li><li><a href="/college/8014.html" target="_blank" title="联创永宣陷股东退出纠纷">联创永宣陷股东退出纠纷</a></li><li><a href="/college/7953.html" target="_blank" title="1元就能投理财子公司产品 需特别关注“理财说明书”">1元就能投理财子公司产品 需特别关注“理财说明书”</a></li><li><a href="/college/7829.html" target="_blank" title="工银理财正式开业 管理理财产品规模超3700亿">工银理财正式开业 管理理财产品规模超3700亿</a></li><li><a href="/college/7818.html" target="_blank" title="已有8家银行理财子公司获批成立 多家开业竞速">已有8家银行理财子公司获批成立 多家开业竞速</a></li><li><a href="/college/7798.html" target="_blank" title="建行旗下理财子公司正式开业运营">建行旗下理财子公司正式开业运营</a></li><li><a href="/college/7585.html" target="_blank" title="网贷返利平台：夏天都来了 你理财了吗？">网贷返利平台：夏天都来了 你理财了吗？</a></li><li><a href="/college/7530.html" target="_blank" title="银行理财这一年:现金管理产品异军突起 净值真假待考">银行理财这一年:现金管理产品异军突起 净值真假待考</a></li><li><a href="/college/7475.html" target="_blank" title="汇添富整改提速 理财60天估值转为市值法">汇添富整改提速 理财60天估值转为市值法</a></li><li><a href="/college/7433.html" target="_blank" title="3月银行理财预期收益率 创两年新低至4.31%">3月银行理财预期收益率 创两年新低至4.31%</a></li>                                            </ul>
                                        </div>
                                    </li>
                                </ul><ul>
                                    <li style="border:0px;padding:0px;">
                                        <div class="post-left">
                                            <div class="left-img">
                                                <a href="/college/6873.html" target="_blank" title="社科院报告：去年个人融资规模接近45万亿元">
                                                                                                    <img src="/data/upload/20181226/5c22eedf579c2.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>社科院报告：去年个人融资规模接近45万亿元</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div><div class="left-img">
                                                <a href="/college/6052.html" target="_blank" title="P2P平台点融获大连金投4000万美元融资">
                                                                                                    <img src="/data/upload/20180807/5b6939d923423.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>P2P平台点融获大连金投4000万美元融资</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div>                                        </div>
                                        <div class="post-right">
                                            <ul>
                                                <li><a href="/college/6873.html" target="_blank" title="社科院报告：去年个人融资规模接近45万亿元">社科院报告：去年个人融资规模接近45万亿元</a></li><li><a href="/college/6052.html" target="_blank" title="P2P平台点融获大连金投4000万美元融资">P2P平台点融获大连金投4000万美元融资</a></li><li><a href="/college/5783.html" target="_blank" title="6月80家P2P平台存问题或停业，发生6例融资事件">6月80家P2P平台存问题或停业，发生6例融资事件</a></li><li><a href="/college/5711.html" target="_blank" title="P2P平台添米财富完成了新一轮千万元融资">P2P平台添米财富完成了新一轮千万元融资</a></li><li><a href="/college/5687.html" target="_blank" title="牛板金再获资本青睐 实缴注册资本提升至6亿元人民币">牛板金再获资本青睐 实缴注册资本提升至6亿元人民币</a></li><li><a href="/college/5586.html" target="_blank" title="蚂蚁金服对外宣布新一轮融资140亿美元">蚂蚁金服对外宣布新一轮融资140亿美元</a></li><li><a href="/college/5520.html" target="_blank" title="懒财金服宣布完成数亿元C轮融资">懒财金服宣布完成数亿元C轮融资</a></li><li><a href="/college/5354.html" target="_blank" title="什马完成3亿元C轮融资 铜板街等机构联合领投">什马完成3亿元C轮融资 铜板街等机构联合领投</a></li><li><a href="/college/5220.html" target="_blank" title="百融金服完成10亿元C轮融资">百融金服完成10亿元C轮融资</a></li><li><a href="/college/4753.html" target="_blank" title="发财猪获A轮融资，融资放疑似准备破产清算">发财猪获A轮融资，融资放疑似准备破产清算</a></li>                                            </ul>
                                        </div>
                                    </li>
                                </ul><ul>
                                    <li style="border:0px;padding:0px;">
                                        <div class="post-left">
                                            <div class="left-img">
                                                <a href="/college/5117.html" target="_blank" title="管家金服欠债不还被纳入法院强制处决，这老赖台子还敢投？">
                                                                                                    <img src="/data/upload/20180403/5ac2f4e11d9c7.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>管家金服欠债不还被纳入法院强制处决，这老赖台子还敢投？</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div><div class="left-img">
                                                <a href="/college/4865.html" target="_blank" title="春节后的“标慌”导致各大平台纷纷降息">
                                                                                                    <img src="/data/upload/20180301/5a979568a4324.png" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>春节后的“标慌”导致各大平台纷纷降息</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div>                                        </div>
                                        <div class="post-right">
                                            <ul>
                                                <li><a href="/college/5117.html" target="_blank" title="管家金服欠债不还被纳入法院强制处决，这老赖台子还敢投？">管家金服欠债不还被纳入法院强制处决，这老赖台子还敢投？</a></li><li><a href="/college/4865.html" target="_blank" title="春节后的“标慌”导致各大平台纷纷降息">春节后的“标慌”导致各大平台纷纷降息</a></li><li><a href="/college/4623.html" target="_blank" title="多彩云刘德来根据客户投资金额时间收取费用">多彩云刘德来根据客户投资金额时间收取费用</a></li><li><a href="/college/4462.html" target="_blank" title="平安集团旗下的“平安好房”兑付困难">平安集团旗下的“平安好房”兑付困难</a></li>                                            </ul>
                                        </div>
                                    </li>
                                </ul><ul>
                                    <li style="border:0px;padding:0px;">
                                        <div class="post-left">
                                            <div class="left-img">
                                                <a href="/college/6688.html" target="_blank" title="深圳从重打击违法P2P 已报送3批恶意逃废债人 ">
                                                                                                    <img src="/data/upload/20181126/5bfba9224b396.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>深圳从重打击违法P2P 已报送3批恶意逃废债人 </h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div><div class="left-img">
                                                <a href="/college/6687.html" target="_blank" title="神秘人“顺其自然”是谁？20年来捐款超1000万">
                                                                                                    <img src="/data/upload/20181126/5bfb6bc064e92.jpg" width="250" height="140" alt="" class="list-img" />
                                                    <div class="text">
                                                        <h3>神秘人“顺其自然”是谁？20年来捐款超1000万</h3>
                                                        <div class="text-bg"></div>
                                                    </div>
                                                </a>
                                            </div>                                        </div>
                                        <div class="post-right">
                                            <ul>
                                                <li><a href="/college/6688.html" target="_blank" title="深圳从重打击违法P2P 已报送3批恶意逃废债人 ">深圳从重打击违法P2P 已报送3批恶意逃废债人 </a></li><li><a href="/college/6687.html" target="_blank" title="神秘人“顺其自然”是谁？20年来捐款超1000万">神秘人“顺其自然”是谁？20年来捐款超1000万</a></li><li><a href="/college/6686.html" target="_blank" title="望洲财富新消息：法院预告拍卖资产 追缴变价后发还受害人">望洲财富新消息：法院预告拍卖资产 追缴变价后发还受害人</a></li><li><a href="/college/6685.html" target="_blank" title="通缉令上“酒托”高颜值女嫌犯火了！网友热议">通缉令上“酒托”高颜值女嫌犯火了！网友热议</a></li><li><a href="/college/6684.html" target="_blank" title="江苏联宝（原扬州宝缘）涉嫌集资诈骗罪、洗钱罪等多项罪名被起诉 ">江苏联宝（原扬州宝缘）涉嫌集资诈骗罪、洗钱罪等多项罪名被起诉 </a></li><li><a href="/college/6683.html" target="_blank" title="早起打卡居然也有骗局！牛轩平台卷款500万跑路！">早起打卡居然也有骗局！牛轩平台卷款500万跑路！</a></li><li><a href="/college/6682.html" target="_blank" title="中普集团因非吸已被立案 已控制帐户409个 归集资金6.09亿元">中普集团因非吸已被立案 已控制帐户409个 归集资金6.09亿元</a></li><li><a href="/college/6681.html" target="_blank" title="收集936万个人的信息已获刑 侵犯公民个人信息罪">收集936万个人的信息已获刑 侵犯公民个人信息罪</a></li><li><a href="/college/6679.html" target="_blank" title="福田分局通报合时代、小零钱、合拍在线进展">福田分局通报合时代、小零钱、合拍在线进展</a></li><li><a href="/college/6678.html" target="_blank" title="饭饭金服查封房产9处 冻结股票12个 冻结保单2份">饭饭金服查封房产9处 冻结股票12个 冻结保单2份</a></li>                                            </ul>
                                        </div>
                                    </li>
                                </ul>                            </div>
                        </div>
                    </div>
                    <!--列表 end-->
                    <!--问答start-->
                    <div class="p2p-list">
                        <div class="p2pBox">
                            <div class="hd">
                                <ul>
                                <style type="text/css">
                                    .p2pBoxulli{padding:0px 16px !important;}
                                </style>
                                <li class="p2pBoxulli"><a href="/ask/askDetail/id/1.html">有问必答</a></li><li class="p2pBoxulli"><a href="/ask/askDetail/id/2.html">投资理财</a></li><li class="p2pBoxulli"><a href="/ask/askDetail/id/3.html">网贷知识</a></li><li class="p2pBoxulli"><a href="/ask/askDetail/id/4.html">贷款百科</a></li><li class="p2pBoxulli"><a href="/ask/askDetail/id/5.html">网贷曝光</a></li><li class="p2pBoxulli"><a href="/ask/askDetail/id/6.html">利息收益</a></li><li class="p2pBoxulli"><a href="/ask/askDetail/id/7.html">平台百科</a></li><li class="p2pBoxulli"><a href="/ask/askDetail/id/8.html">政策法规</a></li>                                </ul>
                            </div>
                            <style type="text/css">
                                .bdulli{width: 50% !important;padding:5px 0px !important;border-bottom:none !important; float: left;}
                                .bdulli>a{font-size: 16px !important;white-space:nowrap;}
                            </style>
                            <div class="bd">
                                <ul>
                                    <li class="bdulli"><a href="/ask/630.html" target="_blank" title="" title="P2P出监管政策，是好事还是坏事?"><!--P2P出监管政策，是好事还是坏事?-->P2P出监管政策，是好事还是坏事?</a></li><li class="bdulli"><a href="/ask/611.html" target="_blank" title="" title="有哪些比余额宝更好的理财方式？"><!--有哪些比余额宝更好的理财方式？-->有哪些比余额宝更好的理财方式？</a></li><li class="bdulli"><a href="/ask/610.html" target="_blank" title="" title="工薪族的投资理财渠道有哪些？"><!--工薪族的投资理财渠道有哪些？-->工薪族的投资理财渠道有哪些？</a></li><li class="bdulli"><a href="/ask/609.html" target="_blank" title="" title="怎样理财收益最大？各种投资理财怎么选？"><!--怎样理财收益最大？各种投资理财怎么选？-->怎样理财收益最大？各种投资理财怎么选？</a></li><li class="bdulli"><a href="/ask/607.html" target="_blank" title="" title="个人投资理财有什么值得推荐的网站？"><!--个人投资理财有什么值得推荐的网站？-->个人投资理财有什么值得推荐的网站？</a></li><li class="bdulli"><a href="/ask/604.html" target="_blank" title="" title="未来p2p网贷应该在哪方面进行创新"><!--未来p2p网贷应该在哪方面进行创新-->未来p2p网贷应该在哪方面进行创新</a></li><li class="bdulli"><a href="/ask/599.html" target="_blank" title="" title="网贷P2P第三方担保靠谱吗？"><!--网贷P2P第三方担保靠谱吗？-->网贷P2P第三方担保靠谱吗？</a></li><li class="bdulli"><a href="/ask/598.html" target="_blank" title="" title="国内最好的P2P理财平台是哪个？"><!--国内最好的P2P理财平台是哪个？-->国内最好的P2P理财平台是哪个？</a></li><li class="bdulli"><a href="/ask/597.html" target="_blank" title="" title="P2P 平台是怎么做风控的？"><!--P2P 平台是怎么做风控的？-->P2P 平台是怎么做风控的？</a></li><li class="bdulli"><a href="/ask/596.html" target="_blank" title="" title="投资p2p真的安全吗？"><!--投资p2p真的安全吗？-->投资p2p真的安全吗？</a></li><li class="bdulli"><a href="/ask/570.html" target="_blank" title="" title="月薪不同的人如何理财？"><!--月薪不同的人如何理财？-->月薪不同的人如何理财？</a></li><li class="bdulli"><a href="/ask/569.html" target="_blank" title="" title="p2p投资有哪些坑？"><!--p2p投资有哪些坑？-->p2p投资有哪些坑？</a></li><li class="bdulli"><a href="/ask/568.html" target="_blank" title="" title="怎么判断一个p2p平台的真实可靠"><!--怎么判断一个p2p平台的真实可靠-->怎么判断一个p2p平台的真实可靠</a></li><li class="bdulli"><a href="/ask/539.html" target="_blank" title="" title=" 是不是太低的利率没人会投资的？多少一个合规的范围？"><!-- 是不是太低的利率没人会投资的？多少一个合规的范围？--> 是不是太低的利率没人会投资的？多少一个合规的…</a></li><li class="bdulli"><a href="/ask/523.html" target="_blank" title="" title="这个可以让朋友注册嘛 是不是也会给我返利"><!--这个可以让朋友注册嘛 是不是也会给我返利-->这个可以让朋友注册嘛 是不是也会给我返利</a></li><li class="bdulli"><a href="/ask/516.html" target="_blank" title="" title="如果实现人生第一个一百万"><!--如果实现人生第一个一百万-->如果实现人生第一个一百万</a></li><li class="bdulli"><a href="/ask/509.html" target="_blank" title="" title=" 什么样的p2p平台够安全？"><!-- 什么样的p2p平台够安全？--> 什么样的p2p平台够安全？</a></li><li class="bdulli"><a href="/ask/508.html" target="_blank" title="" title=" 请问下：选择P2P平台的标准有哪些？请大家帮忙介绍下"><!-- 请问下：选择P2P平台的标准有哪些？请大家帮忙介绍下--> 请问下：选择P2P平台的标准有哪些？请大家帮…</a></li><li class="bdulli"><a href="/ask/495.html" target="_blank" title="" title="该怎样了解判断一家P2P平台的真实安全性？"><!--该怎样了解判断一家P2P平台的真实安全性？-->该怎样了解判断一家P2P平台的真实安全性？</a></li><li class="bdulli"><a href="/ask/486.html" target="_blank" title="" title=" 查询P2P平台背景信息的几种方式？"><!-- 查询P2P平台背景信息的几种方式？--> 查询P2P平台背景信息的几种方式？</a></li>                                  </ul><ul>
                                    <li class="bdulli"><a href="/ask/627.html" target="_blank" title="" title="个人投资如何进行资产配置？"><!--个人投资如何进行资产配置？-->个人投资如何进行资产配置？</a></li><li class="bdulli"><a href="/ask/606.html" target="_blank" title="" title="10万元该如何投资理财？"><!--10万元该如何投资理财？-->10万元该如何投资理财？</a></li><li class="bdulli"><a href="/ask/605.html" target="_blank" title="" title="在保本的前提下，收益最高的投资理财方式有哪些？"><!--在保本的前提下，收益最高的投资理财方式有哪些？-->在保本的前提下，收益最高的投资理财方式有哪些？</a></li><li class="bdulli"><a href="/ask/601.html" target="_blank" title="" title="有哪些值得推荐的网贷平台？"><!--有哪些值得推荐的网贷平台？-->有哪些值得推荐的网贷平台？</a></li><li class="bdulli"><a href="/ask/594.html" target="_blank" title="" title="大家除了p2p网贷，还有没尝试过做其他理财"><!--大家除了p2p网贷，还有没尝试过做其他理财-->大家除了p2p网贷，还有没尝试过做其他理财</a></li><li class="bdulli"><a href="/ask/593.html" target="_blank" title="" title="有没年化15%以上，又比较靠谱的p2p平台推荐？"><!--有没年化15%以上，又比较靠谱的p2p平台推荐？-->有没年化15%以上，又比较靠谱的p2p平台推荐…</a></li><li class="bdulli"><a href="/ask/576.html" target="_blank" title="" title="E周行靠谱吗，为什么返利这么高"><!--E周行靠谱吗，为什么返利这么高-->E周行靠谱吗，为什么返利这么高</a></li><li class="bdulli"><a href="/ask/574.html" target="_blank" title="" title="个人投资理财有哪些注意事项"><!--个人投资理财有哪些注意事项-->个人投资理财有哪些注意事项</a></li><li class="bdulli"><a href="/ask/573.html" target="_blank" title="" title="新手投资理财有哪些入门知识要学？"><!--新手投资理财有哪些入门知识要学？-->新手投资理财有哪些入门知识要学？</a></li><li class="bdulli"><a href="/ask/572.html" target="_blank" title="" title="工资高但是没有存款，如何理财？"><!--工资高但是没有存款，如何理财？-->工资高但是没有存款，如何理财？</a></li><li class="bdulli"><a href="/ask/541.html" target="_blank" title="" title="怎么合理的利用手头资金理财"><!--怎么合理的利用手头资金理财-->怎么合理的利用手头资金理财</a></li><li class="bdulli"><a href="/ask/540.html" target="_blank" title="" title="P2P新手试水，第一次投多少钱比较好"><!--P2P新手试水，第一次投多少钱比较好-->P2P新手试水，第一次投多少钱比较好</a></li><li class="bdulli"><a href="/ask/537.html" target="_blank" title="" title=" P2P投资哪个好，帮忙分析一下呗？"><!-- P2P投资哪个好，帮忙分析一下呗？--> P2P投资哪个好，帮忙分析一下呗？</a></li><li class="bdulli"><a href="/ask/536.html" target="_blank" title="" title=" 网上个人贷款平台可靠吗"><!-- 网上个人贷款平台可靠吗--> 网上个人贷款平台可靠吗</a></li><li class="bdulli"><a href="/ask/499.html" target="_blank" title="" title="乐富支付还款客服电话是多少"><!--乐富支付还款客服电话是多少-->乐富支付还款客服电话是多少</a></li><li class="bdulli"><a href="/ask/496.html" target="_blank" title="" title="明天的超级返利日是哪家平台"><!--明天的超级返利日是哪家平台-->明天的超级返利日是哪家平台</a></li><li class="bdulli"><a href="/ask/484.html" target="_blank" title="" title="最近还有什么活动吗？"><!--最近还有什么活动吗？-->最近还有什么活动吗？</a></li><li class="bdulli"><a href="/ask/482.html" target="_blank" title="" title="现在网贷降息这么严重，怎么能减少损失？"><!--现在网贷降息这么严重，怎么能减少损失？-->现在网贷降息这么严重，怎么能减少损失？</a></li><li class="bdulli"><a href="/ask/480.html" target="_blank" title="" title="最近有没有好的平台可以投一下！"><!--最近有没有好的平台可以投一下！-->最近有没有好的平台可以投一下！</a></li><li class="bdulli"><a href="/ask/478.html" target="_blank" title="" title="降息已成趋势，我们如何在这个时期提高投资收益呢？"><!--降息已成趋势，我们如何在这个时期提高投资收益呢？-->降息已成趋势，我们如何在这个时期提高投资收益呢…</a></li>                                  </ul><ul>
                                    <li class="bdulli"><a href="/ask/618.html" target="_blank" title="" title=" 不把鸡蛋放一个篮子里，做好分散投资吗？"><!-- 不把鸡蛋放一个篮子里，做好分散投资吗？--> 不把鸡蛋放一个篮子里，做好分散投资吗？</a></li><li class="bdulli"><a href="/ask/608.html" target="_blank" title="" title="网贷投资理财的风险在哪里？"><!--网贷投资理财的风险在哪里？-->网贷投资理财的风险在哪里？</a></li><li class="bdulli"><a href="/ask/603.html" target="_blank" title="" title="为什么出现很多p2p网贷平台?"><!--为什么出现很多p2p网贷平台?-->为什么出现很多p2p网贷平台?</a></li><li class="bdulli"><a href="/ask/600.html" target="_blank" title="" title="P2P网贷平台排名情况怎么样？"><!--P2P网贷平台排名情况怎么样？-->P2P网贷平台排名情况怎么样？</a></li><li class="bdulli"><a href="/ask/595.html" target="_blank" title="" title="P2P网贷和银行对比，优劣主要有哪些？"><!--P2P网贷和银行对比，优劣主要有哪些？-->P2P网贷和银行对比，优劣主要有哪些？</a></li><li class="bdulli"><a href="/ask/592.html" target="_blank" title="" title="网贷如何避雷，如何考察平台呢"><!--网贷如何避雷，如何考察平台呢-->网贷如何避雷，如何考察平台呢</a></li><li class="bdulli"><a href="/ask/551.html" target="_blank" title="" title=" 相对其他产品，P2P网贷的优势在哪里？"><!-- 相对其他产品，P2P网贷的优势在哪里？--> 相对其他产品，P2P网贷的优势在哪里？</a></li><li class="bdulli"><a href="/ask/546.html" target="_blank" title="" title=" 新手理财投资要注意什么呀？"><!-- 新手理财投资要注意什么呀？--> 新手理财投资要注意什么呀？</a></li><li class="bdulli"><a href="/ask/538.html" target="_blank" title="" title=" 资深可靠的p2p投资平台都有哪些？"><!-- 资深可靠的p2p投资平台都有哪些？--> 资深可靠的p2p投资平台都有哪些？</a></li><li class="bdulli"><a href="/ask/507.html" target="_blank" title="" title=" P2P年化收益率多少合理"><!-- P2P年化收益率多少合理--> P2P年化收益率多少合理</a></li><li class="bdulli"><a href="/ask/506.html" target="_blank" title="" title=" 经常听到平台说100%本息保障，这承诺靠谱吗？?"><!-- 经常听到平台说100%本息保障，这承诺靠谱吗？?--> 经常听到平台说100%本息保障，这承诺靠谱吗…</a></li><li class="bdulli"><a href="/ask/505.html" target="_blank" title="" title=" 有什么挑选p2p平台的方法？"><!-- 有什么挑选p2p平台的方法？--> 有什么挑选p2p平台的方法？</a></li><li class="bdulli"><a href="/ask/501.html" target="_blank" title="" title=" p2p网贷平台逾期不还款怎么维权？                                    "><!-- p2p网贷平台逾期不还款怎么维权？                                    --> p2p网贷平台逾期不还款怎么维权？     …</a></li><li class="bdulli"><a href="/ask/500.html" target="_blank" title="" title="快钱支付还款客服电话是多少"><!--快钱支付还款客服电话是多少-->快钱支付还款客服电话是多少</a></li><li class="bdulli"><a href="/ask/498.html" target="_blank" title="" title="判断一个平台安全与否都要看那几点？"><!--判断一个平台安全与否都要看那几点？-->判断一个平台安全与否都要看那几点？</a></li><li class="bdulli"><a href="/ask/497.html" target="_blank" title="" title="禁净值标，陆金所、红岭创投何去何从？"><!--禁净值标，陆金所、红岭创投何去何从？-->禁净值标，陆金所、红岭创投何去何从？</a></li><li class="bdulli"><a href="/ask/493.html" target="_blank" title="" title=" P2P平台上的标那些类型的不靠谱？"><!-- P2P平台上的标那些类型的不靠谱？--> P2P平台上的标那些类型的不靠谱？</a></li><li class="bdulli"><a href="/ask/491.html" target="_blank" title="" title="为何最近看到好多平台还是有超过100w的标啊？？？"><!--为何最近看到好多平台还是有超过100w的标啊？？？-->为何最近看到好多平台还是有超过100w的标啊？…</a></li><li class="bdulli"><a href="/ask/489.html" target="_blank" title="" title=" 对于评价P2P平台，第三方资金托管有多重要"><!-- 对于评价P2P平台，第三方资金托管有多重要--> 对于评价P2P平台，第三方资金托管有多重要</a></li><li class="bdulli"><a href="/ask/488.html" target="_blank" title="" title=" 去平台实地考察都要看哪些地方？"><!-- 去平台实地考察都要看哪些地方？--> 去平台实地考察都要看哪些地方？</a></li>                                  </ul><ul>
                                    <li class="bdulli"><a href="/ask/602.html" target="_blank" title="" title="怎么看待P2P网贷？P2P网贷平台发展前景怎么样？"><!--怎么看待P2P网贷？P2P网贷平台发展前景怎么样？-->怎么看待P2P网贷？P2P网贷平台发展前景怎么…</a></li><li class="bdulli"><a href="/ask/565.html" target="_blank" title="" title=" 大型的网贷门户网站除了之家，天眼，网贷互联还有其他的吗"><!-- 大型的网贷门户网站除了之家，天眼，网贷互联还有其他的吗--> 大型的网贷门户网站除了之家，天眼，网贷互联还…</a></li><li class="bdulli"><a href="/ask/562.html" target="_blank" title="" title=" 今后P2P会成为理财主流吗？"><!-- 今后P2P会成为理财主流吗？--> 今后P2P会成为理财主流吗？</a></li><li class="bdulli"><a href="/ask/548.html" target="_blank" title="" title=" 谁能分析一下投资p2p车贷的优缺点？"><!-- 谁能分析一下投资p2p车贷的优缺点？--> 谁能分析一下投资p2p车贷的优缺点？</a></li><li class="bdulli"><a href="/ask/547.html" target="_blank" title="" title="懒丁贷是个什么平台？"><!--懒丁贷是个什么平台？-->懒丁贷是个什么平台？</a></li><li class="bdulli"><a href="/ask/524.html" target="_blank" title="" title="用钱宝还款电话是多少"><!--用钱宝还款电话是多少-->用钱宝还款电话是多少</a></li><li class="bdulli"><a href="/ask/510.html" target="_blank" title="" title=" 最近发现很多人都开始投P2P平台了，大家为什么如此热衷呢？"><!-- 最近发现很多人都开始投P2P平台了，大家为什么如此热衷呢？--> 最近发现很多人都开始投P2P平台了，大家为什…</a></li><li class="bdulli"><a href="/ask/487.html" target="_blank" title="" title=" P2p网贷理财怎么减少投资风险呢啊？"><!-- P2p网贷理财怎么减少投资风险呢啊？--> P2p网贷理财怎么减少投资风险呢啊？</a></li><li class="bdulli"><a href="/ask/443.html" target="_blank" title="" title="理财返利平台是怎样审核P2P平台的？"><!--理财返利平台是怎样审核P2P平台的？-->理财返利平台是怎样审核P2P平台的？</a></li><li class="bdulli"><a href="/ask/437.html" target="_blank" title="" title="钱内助这是个怎么样的平台？"><!--钱内助这是个怎么样的平台？-->钱内助这是个怎么样的平台？</a></li><li class="bdulli"><a href="/ask/433.html" target="_blank" title="" title=" 投资P2P行业，投资人需满足什么样的条件才能进行投资？"><!-- 投资P2P行业，投资人需满足什么样的条件才能进行投资？--> 投资P2P行业，投资人需满足什么样的条件才能…</a></li><li class="bdulli"><a href="/ask/430.html" target="_blank" title="" title=" P2P监管倒计时之后行业会怎么样？"><!-- P2P监管倒计时之后行业会怎么样？--> P2P监管倒计时之后行业会怎么样？</a></li><li class="bdulli"><a href="/ask/370.html" target="_blank" title="" title="爱钱进借款逾期上征信吗？"><!--爱钱进借款逾期上征信吗？-->爱钱进借款逾期上征信吗？</a></li><li class="bdulli"><a href="/ask/358.html" target="_blank" title="" title="通过芝麻信用能贷款的有哪些app？根据芝麻信用能贷多少？"><!--通过芝麻信用能贷款的有哪些app？根据芝麻信用能贷多少？-->通过芝麻信用能贷款的有哪些app？根据芝麻信用…</a></li><li class="bdulli"><a href="/ask/333.html" target="_blank" title="" title="任何东西都可以贷款"><!--任何东西都可以贷款-->任何东西都可以贷款</a></li><li class="bdulli"><a href="/ask/332.html" target="_blank" title="" title="抵押贷款的范围是"><!--抵押贷款的范围是-->抵押贷款的范围是</a></li><li class="bdulli"><a href="/ask/320.html" target="_blank" title="" title="大家有有没有好的贷款渠道？"><!--大家有有没有好的贷款渠道？-->大家有有没有好的贷款渠道？</a></li><li class="bdulli"><a href="/ask/315.html" target="_blank" title="" title="蚂蚁借呗额从1万提升到7.5万这样正常吗？"><!--蚂蚁借呗额从1万提升到7.5万这样正常吗？-->蚂蚁借呗额从1万提升到7.5万这样正常吗？</a></li><li class="bdulli"><a href="/ask/313.html" target="_blank" title="" title="P2P网络贷款支持提前还款吗？"><!--P2P网络贷款支持提前还款吗？-->P2P网络贷款支持提前还款吗？</a></li><li class="bdulli"><a href="/ask/311.html" target="_blank" title="" title="在融360上，贷款金额已经申请下来，但是没有提现运用，会不会影响征信？"><!--在融360上，贷款金额已经申请下来，但是没有提现运用，会不会影响征信？-->在融360上，贷款金额已经申请下来，但是没有提…</a></li>                                  </ul><ul>
                                    <li class="bdulli"><a href="/ask/626.html" target="_blank" title="" title=" 新人理财的投资策略是什么?"><!-- 新人理财的投资策略是什么?--> 新人理财的投资策略是什么?</a></li><li class="bdulli"><a href="/ask/624.html" target="_blank" title="" title=" P2P平台的坏账、逾期都是哪些问题引起的?"><!-- P2P平台的坏账、逾期都是哪些问题引起的?--> P2P平台的坏账、逾期都是哪些问题引起的?</a></li><li class="bdulli"><a href="/ask/623.html" target="_blank" title="" title=" 网贷投资理财，投资人到底应该怎么投？"><!-- 网贷投资理财，投资人到底应该怎么投？--> 网贷投资理财，投资人到底应该怎么投？</a></li><li class="bdulli"><a href="/ask/622.html" target="_blank" title="" title=" 投资p2p平台的成功要素有哪些？"><!-- 投资p2p平台的成功要素有哪些？--> 投资p2p平台的成功要素有哪些？</a></li><li class="bdulli"><a href="/ask/621.html" target="_blank" title="" title=" 正规的P2P网贷平台要去哪里找？"><!-- 正规的P2P网贷平台要去哪里找？--> 正规的P2P网贷平台要去哪里找？</a></li><li class="bdulli"><a href="/ask/620.html" target="_blank" title="" title=" p2p网贷行业，是不是哪个公司投放的广 告越多越安全的？"><!-- p2p网贷行业，是不是哪个公司投放的广 告越多越安全的？--> p2p网贷行业，是不是哪个公司投放的广 告越…</a></li><li class="bdulli"><a href="/ask/619.html" target="_blank" title="" title=" P2P网贷理财相对其流动性好吗？"><!-- P2P网贷理财相对其流动性好吗？--> P2P网贷理财相对其流动性好吗？</a></li><li class="bdulli"><a href="/ask/617.html" target="_blank" title="" title=" 哪种理财产品理财方式更受欢迎？"><!-- 哪种理财产品理财方式更受欢迎？--> 哪种理财产品理财方式更受欢迎？</a></li><li class="bdulli"><a href="/ask/616.html" target="_blank" title="" title=" 国家支持监管细则的出台的初衷是什么？"><!-- 国家支持监管细则的出台的初衷是什么？--> 国家支持监管细则的出台的初衷是什么？</a></li><li class="bdulli"><a href="/ask/615.html" target="_blank" title="" title=" 现在“价低安全”的个人住房按揭贷款有望迎来大幅增长是吗？"><!-- 现在“价低安全”的个人住房按揭贷款有望迎来大幅增长是吗？--> 现在“价低安全”的个人住房按揭贷款有望迎来大…</a></li><li class="bdulli"><a href="/ask/614.html" target="_blank" title="" title=" 什么是P2P平台的风投控制？"><!-- 什么是P2P平台的风投控制？--> 什么是P2P平台的风投控制？</a></li><li class="bdulli"><a href="/ask/613.html" target="_blank" title="" title=" 请问现在都有哪些银行取消了P2P交易接口？"><!-- 请问现在都有哪些银行取消了P2P交易接口？--> 请问现在都有哪些银行取消了P2P交易接口？</a></li><li class="bdulli"><a href="/ask/612.html" target="_blank" title="" title=" 现在互联网投资理财最受欢迎是吗？"><!-- 现在互联网投资理财最受欢迎是吗？--> 现在互联网投资理财最受欢迎是吗？</a></li><li class="bdulli"><a href="/ask/591.html" target="_blank" title="" title=" 请教一下，平台的利率是谁决定的？"><!-- 请教一下，平台的利率是谁决定的？--> 请教一下，平台的利率是谁决定的？</a></li><li class="bdulli"><a href="/ask/590.html" target="_blank" title="" title=" P2P债权转让靠谱吗？有啥风险？"><!-- P2P债权转让靠谱吗？有啥风险？--> P2P债权转让靠谱吗？有啥风险？</a></li><li class="bdulli"><a href="/ask/589.html" target="_blank" title="" title=" 如何预测P2P平台的风险？"><!-- 如何预测P2P平台的风险？--> 如何预测P2P平台的风险？</a></li><li class="bdulli"><a href="/ask/588.html" target="_blank" title="" title=" p2p借贷发展史是怎么回事？"><!-- p2p借贷发展史是怎么回事？--> p2p借贷发展史是怎么回事？</a></li><li class="bdulli"><a href="/ask/587.html" target="_blank" title="" title=" 银行利率持续下滑，资金难道都流向P2P了？"><!-- 银行利率持续下滑，资金难道都流向P2P了？--> 银行利率持续下滑，资金难道都流向P2P了？</a></li><li class="bdulli"><a href="/ask/586.html" target="_blank" title="" title=" 网络借贷行业标准是什么？"><!-- 网络借贷行业标准是什么？--> 网络借贷行业标准是什么？</a></li><li class="bdulli"><a href="/ask/582.html" target="_blank" title="" title=" 选择P2P网贷，哪种还款方式最好？"><!-- 选择P2P网贷，哪种还款方式最好？--> 选择P2P网贷，哪种还款方式最好？</a></li>                                  </ul><ul>
                                    <li class="bdulli"><a href="/ask/544.html" target="_blank" title="" title=" 短期理财比较好还是活期理财比较好呢？"><!-- 短期理财比较好还是活期理财比较好呢？--> 短期理财比较好还是活期理财比较好呢？</a></li><li class="bdulli"><a href="/ask/475.html" target="_blank" title="" title="银行理财产品利率提高，p2p行业的利率将会有什么变化.？"><!--银行理财产品利率提高，p2p行业的利率将会有什么变化.？-->银行理财产品利率提高，p2p行业的利率将会有什…</a></li><li class="bdulli"><a href="/ask/426.html" target="_blank" title="" title=" 有本息保障承诺的平台是否就会更安全？"><!-- 有本息保障承诺的平台是否就会更安全？--> 有本息保障承诺的平台是否就会更安全？</a></li><li class="bdulli"><a href="/ask/420.html" target="_blank" title="" title="大家选择P2P平台是看重收益还是背景？有没有平台可以推荐的？"><!--大家选择P2P平台是看重收益还是背景？有没有平台可以推荐的？-->大家选择P2P平台是看重收益还是背景？有没有平…</a></li><li class="bdulli"><a href="/ask/380.html" target="_blank" title="" title="等额本息和等额本金的区别是什么啊？"><!--等额本息和等额本金的区别是什么啊？-->等额本息和等额本金的区别是什么啊？</a></li><li class="bdulli"><a href="/ask/372.html" target="_blank" title="" title="凤凰金融利率多少？"><!--凤凰金融利率多少？-->凤凰金融利率多少？</a></li><li class="bdulli"><a href="/ask/359.html" target="_blank" title="" title="在e周行平台直投的利率为什么没有在银桥网上高？"><!--在e周行平台直投的利率为什么没有在银桥网上高？-->在e周行平台直投的利率为什么没有在银桥网上高？</a></li><li class="bdulli"><a href="/ask/351.html" target="_blank" title="" title="大家的P2P利率都在多少啊？"><!--大家的P2P利率都在多少啊？-->大家的P2P利率都在多少啊？</a></li><li class="bdulli"><a href="/ask/324.html" target="_blank" title="" title="P2P额外加息好，还是在原本的利息上加息好，哪个安全?"><!--P2P额外加息好，还是在原本的利息上加息好，哪个安全?-->P2P额外加息好，还是在原本的利息上加息好，哪…</a></li><li class="bdulli"><a href="/ask/323.html" target="_blank" title="" title="银行结构型理财产品7%收益，普通人可以买吗"><!--银行结构型理财产品7%收益，普通人可以买吗-->银行结构型理财产品7%收益，普通人可以买吗</a></li><li class="bdulli"><a href="/ask/303.html" target="_blank" title="" title="P2P平台的利率都是怎么计算的？"><!--P2P平台的利率都是怎么计算的？-->P2P平台的利率都是怎么计算的？</a></li><li class="bdulli"><a href="/ask/302.html" target="_blank" title="" title="什么样的原因导致P2P降息？"><!--什么样的原因导致P2P降息？-->什么样的原因导致P2P降息？</a></li><li class="bdulli"><a href="/ask/273.html" target="_blank" title="" title="农业银行的抵押房产的利息是怎么算的？"><!--农业银行的抵押房产的利息是怎么算的？-->农业银行的抵押房产的利息是怎么算的？</a></li><li class="bdulli"><a href="/ask/272.html" target="_blank" title="" title="农业银行的汽车抵押贷款的利息怎么算的？"><!--农业银行的汽车抵押贷款的利息怎么算的？-->农业银行的汽车抵押贷款的利息怎么算的？</a></li><li class="bdulli"><a href="/ask/269.html" target="_blank" title="" title="农业银行个人小额贷款申请条件有哪些？利率是多少"><!--农业银行个人小额贷款申请条件有哪些？利率是多少-->农业银行个人小额贷款申请条件有哪些？利率是多少</a></li><li class="bdulli"><a href="/ask/264.html" target="_blank" title="" title="网贷理财中网贷利息会不会高出本金呢？"><!--网贷理财中网贷利息会不会高出本金呢？-->网贷理财中网贷利息会不会高出本金呢？</a></li><li class="bdulli"><a href="/ask/262.html" target="_blank" title="" title="大家现阶段的收益大约是多少啊？"><!--大家现阶段的收益大约是多少啊？-->大家现阶段的收益大约是多少啊？</a></li><li class="bdulli"><a href="/ask/191.html" target="_blank" title="" title="我朋友介绍我去3M金融圈，这个收益很高！"><!--我朋友介绍我去3M金融圈，这个收益很高！-->我朋友介绍我去3M金融圈，这个收益很高！</a></li><li class="bdulli"><a href="/ask/176.html" target="_blank" title="" title="P2P理财、银行活期存款、余额宝这三种利率有什么样的区别？"><!--P2P理财、银行活期存款、余额宝这三种利率有什么样的区别？-->P2P理财、银行活期存款、余额宝这三种利率有什…</a></li><li class="bdulli"><a href="/ask/170.html" target="_blank" title="" title="农业银行的活期利息怎么算的，比我我有20万，一个月后收益是多？"><!--农业银行的活期利息怎么算的，比我我有20万，一个月后收益是多？-->农业银行的活期利息怎么算的，比我我有20万，一…</a></li>                                  </ul><ul>
                                    <li class="bdulli"><a href="/ask/585.html" target="_blank" title="" title="好车贷，这个平台可以吗？"><!--好车贷，这个平台可以吗？-->好车贷，这个平台可以吗？</a></li><li class="bdulli"><a href="/ask/584.html" target="_blank" title="" title="巨人理财这个平台如何？？"><!--巨人理财这个平台如何？？-->巨人理财这个平台如何？？</a></li><li class="bdulli"><a href="/ask/567.html" target="_blank" title="" title=" 下半年想投资一些收益中上、银行存管等齐备的平台有何推荐？"><!-- 下半年想投资一些收益中上、银行存管等齐备的平台有何推荐？--> 下半年想投资一些收益中上、银行存管等齐备的平…</a></li><li class="bdulli"><a href="/ask/556.html" target="_blank" title="" title="融金所好像不错，我关注好久了，大家觉得如何？？"><!--融金所好像不错，我关注好久了，大家觉得如何？？-->融金所好像不错，我关注好久了，大家觉得如何？？</a></li><li class="bdulli"><a href="/ask/555.html" target="_blank" title="" title="金陵贷最近怎么样啦，可以不？"><!--金陵贷最近怎么样啦，可以不？-->金陵贷最近怎么样啦，可以不？</a></li><li class="bdulli"><a href="/ask/542.html" target="_blank" title="" title=" 有哪些用户体验比较好的理财平台，求推荐。"><!-- 有哪些用户体验比较好的理财平台，求推荐。--> 有哪些用户体验比较好的理财平台，求推荐。</a></li><li class="bdulli"><a href="/ask/522.html" target="_blank" title="" title="大家推荐一个安全稳定，收益合理的小平台"><!--大家推荐一个安全稳定，收益合理的小平台-->大家推荐一个安全稳定，收益合理的小平台</a></li><li class="bdulli"><a href="/ask/515.html" target="_blank" title="" title="监管日益严格，诸多P2P平台如何在不降低风控水平的前提下控制成本"><!--监管日益严格，诸多P2P平台如何在不降低风控水平的前提下控制成本-->监管日益严格，诸多P2P平台如何在不降低风控水…</a></li><li class="bdulli"><a href="/ask/514.html" target="_blank" title="" title="最近看了下易通贷发现还不错，大家觉得呢？"><!--最近看了下易通贷发现还不错，大家觉得呢？-->最近看了下易通贷发现还不错，大家觉得呢？</a></li><li class="bdulli"><a href="/ask/513.html" target="_blank" title="" title="融金所这个平台靠谱不？"><!--融金所这个平台靠谱不？-->融金所这个平台靠谱不？</a></li><li class="bdulli"><a href="/ask/504.html" target="_blank" title="" title="链金所这个平台如何？"><!--链金所这个平台如何？-->链金所这个平台如何？</a></li><li class="bdulli"><a href="/ask/502.html" target="_blank" title="" title="花果金融这个平台怎么样？                       "><!--花果金融这个平台怎么样？                       -->花果金融这个平台怎么样？           …</a></li><li class="bdulli"><a href="/ask/490.html" target="_blank" title="" title="小树时代这个平台如何呢？"><!--小树时代这个平台如何呢？-->小树时代这个平台如何呢？</a></li><li class="bdulli"><a href="/ask/485.html" target="_blank" title="" title="链金所这个平台怎么样？"><!--链金所这个平台怎么样？-->链金所这个平台怎么样？</a></li><li class="bdulli"><a href="/ask/481.html" target="_blank" title="" title="拍拍贷的借款利息高吗？         "><!--拍拍贷的借款利息高吗？         -->拍拍贷的借款利息高吗？         </a></li><li class="bdulli"><a href="/ask/450.html" target="_blank" title="" title="不做网贷的红岭接下来要干啥？"><!--不做网贷的红岭接下来要干啥？-->不做网贷的红岭接下来要干啥？</a></li><li class="bdulli"><a href="/ask/441.html" target="_blank" title="" title="友金所怎么样？靠谱不？？？"><!--友金所怎么样？靠谱不？？？-->友金所怎么样？靠谱不？？？</a></li><li class="bdulli"><a href="/ask/440.html" target="_blank" title="" title="友金所这个平台怎么样？"><!--友金所这个平台怎么样？-->友金所这个平台怎么样？</a></li><li class="bdulli"><a href="/ask/428.html" target="_blank" title="" title="赶钱网这个新加盟的平台怎么样？"><!--赶钱网这个新加盟的平台怎么样？-->赶钱网这个新加盟的平台怎么样？</a></li><li class="bdulli"><a href="/ask/416.html" target="_blank" title="" title="金领贷怎么样？"><!--金领贷怎么样？-->金领贷怎么样？</a></li>                                  </ul><ul>
                                    <li class="bdulli"><a href="/ask/629.html" target="_blank" title="" title="大变局前夜，新监管政策下P2P将走向何方？"><!--大变局前夜，新监管政策下P2P将走向何方？-->大变局前夜，新监管政策下P2P将走向何方？</a></li><li class="bdulli"><a href="/ask/628.html" target="_blank" title="" title="请教一下稳健的货币政策和金融去杠杆是否矛盾？"><!--请教一下稳健的货币政策和金融去杠杆是否矛盾？-->请教一下稳健的货币政策和金融去杠杆是否矛盾？</a></li><li class="bdulli"><a href="/ask/494.html" target="_blank" title="" title="监管为什么要求平台存管属地化？"><!--监管为什么要求平台存管属地化？-->监管为什么要求平台存管属地化？</a></li><li class="bdulli"><a href="/ask/492.html" target="_blank" title="" title=" 目前还没有存管的平台该何去何从？"><!-- 目前还没有存管的平台该何去何从？--> 目前还没有存管的平台该何去何从？</a></li><li class="bdulli"><a href="/ask/471.html" target="_blank" title="" title=" 现在监管细则是怎样的？风险准备金这块是什么意见？"><!-- 现在监管细则是怎样的？风险准备金这块是什么意见？--> 现在监管细则是怎样的？风险准备金这块是什么意…</a></li><li class="bdulli"><a href="/ask/421.html" target="_blank" title="" title=" p2p平台的风险备用金能保障本金安全吗？"><!-- p2p平台的风险备用金能保障本金安全吗？--> p2p平台的风险备用金能保障本金安全吗？</a></li><li class="bdulli"><a href="/ask/411.html" target="_blank" title="" title="网贷所有类型中是不是车贷的接受度最高！？"><!--网贷所有类型中是不是车贷的接受度最高！？-->网贷所有类型中是不是车贷的接受度最高！？</a></li><li class="bdulli"><a href="/ask/406.html" target="_blank" title="" title="独家|惊天内幕？曝同一老板炮制多家跑路平台"><!--独家|惊天内幕？曝同一老板炮制多家跑路平台-->独家|惊天内幕？曝同一老板炮制多家跑路平台</a></li><li class="bdulli"><a href="/ask/398.html" target="_blank" title="" title="广东要求P2P平台禁止个人间债权转让！有病吗？"><!--广东要求P2P平台禁止个人间债权转让！有病吗？-->广东要求P2P平台禁止个人间债权转让！有病吗？</a></li><li class="bdulli"><a href="/ask/375.html" target="_blank" title="" title="年化7%银行理财是不是真的？"><!--年化7%银行理财是不是真的？-->年化7%银行理财是不是真的？</a></li><li class="bdulli"><a href="/ask/362.html" target="_blank" title="" title="穿透性监管中的穿透性是指的什么？"><!--穿透性监管中的穿透性是指的什么？-->穿透性监管中的穿透性是指的什么？</a></li><li class="bdulli"><a href="/ask/321.html" target="_blank" title="" title="“银行存管归属化”到底还能不能落实啊？"><!--“银行存管归属化”到底还能不能落实啊？-->“银行存管归属化”到底还能不能落实啊？</a></li><li class="bdulli"><a href="/ask/318.html" target="_blank" title="" title="民间借贷的利率无上限吗？高出多少将不受法律的保护？"><!--民间借贷的利率无上限吗？高出多少将不受法律的保护？-->民间借贷的利率无上限吗？高出多少将不受法律的保…</a></li><li class="bdulli"><a href="/ask/304.html" target="_blank" title="" title="P2P平台的年化利率超过40%合规吗？法律管的到吗？"><!--P2P平台的年化利率超过40%合规吗？法律管的到吗？-->P2P平台的年化利率超过40%合规吗？法律管的…</a></li><li class="bdulli"><a href="/ask/279.html" target="_blank" title="" title="《关于促进互联网金融健康发展的指导意见》对小额贷款的说明"><!--《关于促进互联网金融健康发展的指导意见》对小额贷款的说明-->《关于促进互联网金融健康发展的指导意见》对小额…</a></li><li class="bdulli"><a href="/ask/278.html" target="_blank" title="" title="北京地区网贷监管细则全文中最核心的内容强调了什么？"><!--北京地区网贷监管细则全文中最核心的内容强调了什么？-->北京地区网贷监管细则全文中最核心的内容强调了什…</a></li><li class="bdulli"><a href="/ask/267.html" target="_blank" title="" title="个人民间贷在法律上承认吗？"><!--个人民间贷在法律上承认吗？-->个人民间贷在法律上承认吗？</a></li><li class="bdulli"><a href="/ask/252.html" target="_blank" title="" title="P2P签订的电子合同协议有法律效力吗？"><!--P2P签订的电子合同协议有法律效力吗？-->P2P签订的电子合同协议有法律效力吗？</a></li><li class="bdulli"><a href="/ask/239.html" target="_blank" title="" title="《通知》中提及的升级校园贷治理的工作要点有哪些？"><!--《通知》中提及的升级校园贷治理的工作要点有哪些？-->《通知》中提及的升级校园贷治理的工作要点有哪些…</a></li><li class="bdulli"><a href="/ask/232.html" target="_blank" title="" title="“自融”后，把资金用在其他地方，法律对这样的情况是什么态度？"><!--“自融”后，把资金用在其他地方，法律对这样的情况是什么态度？-->“自融”后，把资金用在其他地方，法律对这样的情…</a></li>                                  </ul>                            </div>
                        </div>
                    </div>
                    <!--问答 end-->
                </div>
                <link href="/public/simpleboot/font-awesome/4.4.0/css/font-awesome.min.css"  rel="stylesheet" type="text/css">
<div class="college-right">
    <div class="label">
        <div class="college-title"><span class="title-border"></span>热门标签</div>
        <div class="label-list">
            <ul>
                <li><a href="/college/tag/1.html" target="_blank"><i class="fa fa-p2p fa-2x" style="width:50px;margin-top: 6px"></i>P2P</a></li><li><a href="/college/tag/2.html" target="_blank"><i class="fa fa-jinrong fa-2x" style="width:50px;margin-top: 6px"></i>金融</a></li><li><a href="/college/tag/3.html" target="_blank"><i class="fa fa-pintai fa-2x" style="width:50px;margin-top: 6px"></i>平台</a></li><li><a href="/college/tag/4.html" target="_blank"><i class="fa fa-shuju fa-2x" style="width:50px;margin-top: 6px"></i>数据</a></li><li><a href="/college/tag/5.html" target="_blank"><i class="fa fa-rongzi fa-2x" style="width:50px;margin-top: 6px"></i>融资</a></li><li><a href="/college/tag/6.html" target="_blank"><i class="fa fa-zhengce fa-2x" style="width:50px;margin-top: 6px"></i>政策</a></li><li><a href="/college/tag/7.html" target="_blank"><i class="fa fa-licai fa-2x" style="width:50px;margin-top: 6px"></i>理财</a></li><li><a href="/college/tag/8.html" target="_blank"><i class="fa fa-lilv fa-2x" style="width:50px;margin-top: 6px"></i>利率</a></li>            </ul>
        </div>
    </div>
    <div class="label">
        <div class="college-title title-bottom"><span class="title-border"></span>热门排行</div>
        <div class="label-ranking">
            <!--切换-->
            <div class="rankingBox">
                <div class="hd">
                    <ul>
                        <li class="on">月</li>
                        <li>周</li>
                        <li>日</li>
                        
                    </ul>
                </div>
                <div class="bd">
                    <ul>
                        <li>
                            <span class="date"></span>
                            <a href="/college/8015.html" title="信而富发布6月平台回款公告 并推出“共赢计划”" target="_blank">信而富发布6月平台回款公告 并推出“共赢计划”</a>
                        </li><li>
                            <span class="date"></span>
                            <a href="/college/8016.html" title="恒富金融涉嫌非吸被立案 限制涉案人员出入境" target="_blank">恒富金融涉嫌非吸被立案 限制涉案人员出入境</a>
                        </li><li>
                            <span class="date"></span>
                            <a href="/college/8017.html" title="苏宁易购完成剥离旗下子公司苏宁小店 预计增加净利34.28亿" target="_blank">苏宁易购完成剥离旗下子公司苏宁小店 预计增加净</a>
                        </li><li>
                            <span class="date"></span>
                            <a href="/college/8018.html" title="今年以来8家信托公司领14张罚单 信托业严监管势头不减" target="_blank">今年以来8家信托公司领14张罚单 信托业严监管</a>
                        </li><li>
                            <span class="date"></span>
                            <a href="/college/8019.html" title="招行钱端事件立案背后：实控人陈强爆雷前涉诉9.22亿" target="_blank">招行钱端事件立案背后：实控人陈强爆雷前涉诉9.</a>
                        </li><li>
                            <span class="date"></span>
                            <a href="/college/8020.html" title="6月网贷月报：持续保持“三降”趋势 正常运营平台数跌破900家" target="_blank">6月网贷月报：持续保持“三降”趋势 正常运营平</a>
                        </li>                    </ul>
                    <ul style="display: none;">
                        <li>
                            <span class="date"></span>
                            <a href="/college/8073.html" title="网信发布最新通知：延迟原定下周一的投资人见面会" target="_blank">网信发布最新通知：延迟原定下周一的投资人见面会</a>
                        </li><li>
                            <span class="date"></span>
                            <a href="/college/8074.html" title=" 德银耗资74亿欧元整改：退出股票交易业务 拟裁员2万人" target="_blank"> 德银耗资74亿欧元整改：退出股票交易业务 拟</a>
                        </li><li>
                            <span class="date"></span>
                            <a href="/college/8075.html" title="排放造假被罚1.7亿元 江淮汽车大跌逾6.07％" target="_blank">排放造假被罚1.7亿元 江淮汽车大跌逾6.07</a>
                        </li><li>
                            <span class="date"></span>
                            <a href="/college/8076.html" title="宜贷网发布致歉信:兑付不尽人意 需要更多时间" target="_blank">宜贷网发布致歉信:兑付不尽人意 需要更多时间</a>
                        </li><li>
                            <span class="date"></span>
                            <a href="/college/8077.html" title="拿去花资产暴涨背后保理ABS不良大增 携程能否复制花呗小贷？" target="_blank">拿去花资产暴涨背后保理ABS不良大增 携程能否</a>
                        </li><li>
                            <span class="date"></span>
                            <a href="/college/8078.html" title="交行发售理财子公司产品：1元起投 股市投资比例最高20％" target="_blank">交行发售理财子公司产品：1元起投 股市投资比例</a>
                        </li>                    </ul>
                    <ul style="display: none;">
                        <li>
                            <span class="date"></span>
                            <a href="/college/8073.html" title="网信发布最新通知：延迟原定下周一的投资人见面会" target="_blank">网信发布最新通知：延迟原定下周一的投资人见面会</a>
                        </li><li>
                            <span class="date"></span>
                            <a href="/college/8074.html" title=" 德银耗资74亿欧元整改：退出股票交易业务 拟裁员2万人" target="_blank"> 德银耗资74亿欧元整改：退出股票交易业务 拟</a>
                        </li><li>
                            <span class="date"></span>
                            <a href="/college/8075.html" title="排放造假被罚1.7亿元 江淮汽车大跌逾6.07％" target="_blank">排放造假被罚1.7亿元 江淮汽车大跌逾6.07</a>
                        </li><li>
                            <span class="date"></span>
                            <a href="/college/8076.html" title="宜贷网发布致歉信:兑付不尽人意 需要更多时间" target="_blank">宜贷网发布致歉信:兑付不尽人意 需要更多时间</a>
                        </li><li>
                            <span class="date"></span>
                            <a href="/college/8077.html" title="拿去花资产暴涨背后保理ABS不良大增 携程能否复制花呗小贷？" target="_blank">拿去花资产暴涨背后保理ABS不良大增 携程能否</a>
                        </li><li>
                            <span class="date"></span>
                            <a href="/college/8078.html" title="交行发售理财子公司产品：1元起投 股市投资比例最高20％" target="_blank">交行发售理财子公司产品：1元起投 股市投资比例</a>
                        </li>                    </ul>
                    
                </div>
            </div>
            <!--切换-->
        </div>
    </div>
<!--     <div class="label">
        <div class="college-title"><span class="title-border"></span>媒体报道<a href="#" target="_blank">更多+</a></div>
        <div class="label-report">
            <div class="report-heat">
                <img src="/assets/images/test/img-128-86.png" width="128" height="86" alt="" class="list-img" />
                <div class="list-det">
                    <h3><a href="#" target="_blank">p2p行业的投资者</a></h3>
                    <p class="details">虽然近两年来p2p平台暴雷事件影响非常恶劣，但是其实...</p>
                </div>
            </div>
            <div class="headlines-list clearfix">
                <ul>
                    <li><a href="http://www.SuperSlide2.com" target="_blank">44家互金独角兽估值1287亿美元 资本最为活跃</a></li>
                    <li><a href="http://www.SuperSlide2.com" target="_blank">停业已非丑闻 P2P网贷行业逐渐趋于良性</a></li>
                    <li><a href="http://www.SuperSlide2.com" target="_blank">米缸金融获CFCA电子签章认证 加固电子合同安全</a>
                    </li>
                    <li><a href="http://www.SuperSlide2.com" target="_blank">中国打破了世界软件巨头规则</a></li>
                    <li><a href="http://www.SuperSlide2.com" target="_blank">中国打破了世界软件巨头规则</a></li>
                    <li><a href="http://www.SuperSlide2.com" target="_blank">中国打破了世界软件巨头规则</a></li>
                </ul>
            </div>
        </div>
    </div> -->
</div>
            </div>
        </div>
    </div>
    
<!--底部通用-->
<style>
footer ul.clearfix{
	margin: 0 auto;
	position: relative;
	left: 49%;
}
footer ul.clearfix li {
	float: left;
	width: 140px;
	margin: 0 10px 30px;
}
</style>
	<footer>
		<div style="background:#FFF!important;" class="clearfix">
		<div class="footer_center clearfix">
			<div class="footer_left">
				<dl>
					<dt><i class="icon icon_about mr12"></i>关于我们</dt>
					<dd><a href="/aboutus/index.html" target="_blank">了解我们</a></dd>
					<dd><a href="/aboutus/advertise.html" target="_blank">加入我们</a></dd>
					<dd><a href="/aboutus/contactus.html" target="_blank">联系我们</a></dd>
				</dl>
				<dl>
					<dt><i class="icon icon_help mr12"></i>服务工具箱</dt>
					<dd><a href="/faq.html" target="_blank">新手上路</a></dd>
					<dd><a href="/aboutus/agreement.html" target="_blank">服务协议</a></dd>
					<dd><a href="/aboutus/scafina.html" target="_blank">收益计算器</a></dd>
				</dl>
				<dl>
					<dt><i class="icon icon_pro mr12"></i>产品服务</dt>
					<dd><a href="/witkey.html" target="_blank">首投返利</a></dd>
					<dd><a href="/rebate.html" target="_blank">复投专享</a></dd>
				</dl>
				<dl style="border-left: solid 1px #f3f3f3; border-right: solid 1px #f3f3f3; margin-right: 0px; padding-right: 100px; padding-left: 80px;">
					<dt><i class="icon icon_see mr12"></i></dt>
					<dd style="height:95px;"><img src="/themes/simplebootx/Public/images/weixin.jpg" height="100" width="100"></dd>
					<dd><a href="javascript:;">微信公众号</a></dd>
				</dl>
			</div>
			<div class="footer_rig">
				<dl>
					<dt><i class="icon icon_service"></i>服务支持</dt>
					<dd class="fz30 color2">0530-5342994</dd>
					<dd class="mt10">送钱网QQ群：<a target="_blank" href="http://shang.qq.com/wpa/qunwpa?idkey=7456307f5236a04aaad6b05abe978d1dd1cef4b6b29e4f5bd9f3819021267fbd" title="点击申请加群">669600739</a></dd>
					<dd class="mt10">服务时间：9:00-21:00</dd>
					<dd class="mt12"><a href="http://wpa.qq.com/msgrd?v=3&uin=531247347&site=qq&menu=yes">意见反馈</a></dd>
				</dl>
			</div>
		</div>
		<div class="footer_copyright">
			<div class="copyright"><p>运营中心：山东省菏泽市东关街道华瑞花园4号楼2单元</p>
            <p>
                Copyright ? 2013-2016 送钱网 HJCYH.com  山东银灯网络科技有限公司 Inc. All Rights Reserved.&nbsp;&nbsp;&nbsp;<script src="https://s95.cnzz.com/z_stat.php?id=1261038503&web_id=1261038503" language="JavaScript"></script>
<!--                //百度统计-->
                <script>
var _hmt = _hmt || [];
(function() {
  var hm = document.createElement("script");
  hm.src = "https://hm.baidu.com/hm.js?9ec9d849345216f88f7aebef9831ae76";
  var s = document.getElementsByTagName("script")[0]; 
  s.parentNode.insertBefore(hm, s);
})();
</script>                <!--百度统计-->
                &nbsp;&nbsp;&nbsp;<a href="http://www.beian.miit.gov.cn">沪ICP备17048089号-3</a> 
            </p>
            <p>技术支持：<a href="http://www.yintiao.net/" target="_blank">银条软件</a></p>
                        </div>
		</div>
	</div>
	<div style="width:1200px;background:#FFF;">
            <ul class="clearfix">
            <li>
                <a target="_blank" href="https://v.pinpaibao.com.cn/cert/site/?site=www.hjcyh.com&at=realname" ><img src="https://static.anquan.org/static/outer/image/sm_124x47.png"></img></a>
            </li>
                 <li>
                    <a target="_blank" href="https://v.pinpaibao.com.cn/cert/site/?site=www.hjcyh.com&at=realname" ><img src="https://static.anquan.org/static/outer/image/aqkx_124x47.png"></img></a>
                </li>
                
            </ul>
        </div>

	</footer>
        <script src="/themes/simplebootx/Public/js/plugins/flexslider/jquery.flexslider.js"></script>	
        <script src="/themes/simplebootx/Public/js/plugins/jquery.SuperSlide.2.1.1.js"></script>
        <script src="/themes/simplebootx/Public/js/college/college.js"></script>
        <script src="/themes/simplebootx/Public/js/plugins/utility.js"></script>
        <link src="/themes/simplebootx/Public/js/plugins/layer/skin/layer.css" />
        <script src="/themes/simplebootx/Public/js/plugins/layer/layer.js"></script>
		<script>(function(){
var src = (document.location.protocol == "http:") ? "http://js.passport.qihucdn.com/11.0.1.js?5adbee5cd75c42ca22112ee98b1eb516":"https://jspassport.ssl.qhimg.com/11.0.1.js?5adbee5cd75c42ca22112ee98b1eb516";
document.write('<script src="' + src + '" id="sozz"><\/script>');
})();
</script>
<style type="text/css">
#nav_activity{height:260px;width:100%;position:fixed;left:0;bottom:0;opacity:0;z-index:100}
#nav_activity  img{width:100%;}
#nav_activity  #nav_activity-exit{height:30px;width:30px;cursor:pointer;position:absolute;right:40px;top:30px;background:url(/themes/simplebootx/Public/images/public/ad_close.png) no-repeat;transition:transform .5s }
#nav_activity #nav_activity-exit:hover{
    transform:rotate(200deg);
    -ms-transform:rotate(200deg);     /* IE 9 */
    -moz-transform:rotate(200deg);    /* Firefox */
    -webkit-transform:rotate(200deg); /* Safari 和 Chrome */
    -o-transform:rotate(200deg);  /* Opera */
}
</style>
<script type="text/javascript">
  function cloase_ad()
  {
      document.getElementById('nav_activity').style.display="none";
      return false;
  }
</script>
<!-- <div id="nav_activity" style="transition-delay: 0.5s; margin-left: 0px; opacity: 1; display: block;">
    <a href='http://www.hjcyh.com/notice/100.html'><img src='/data/upload/20180609/5b1b8b98a3039.jpg' alt='' width='1920' height='246' /></a>    <div id="nav_activity-exit" onclick="return cloase_ad()"></div>
 </div> -->
</body></html>
<!--底部通用结束-->
    <script src="/themes/simplebootx/Public/js/college/college.js"></script>
`

// innerHTML returns the HTML content (inner HTML) of an element.
func InnerHTML(node *html.Node) string {
	var err error
	var buffer bytes.Buffer

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		err = html.Render(&buffer, child)
		if err != nil {
			return ""
		}
	}

	return strings.TrimSpace(buffer.String())
}
func TestNewParser(t *testing.T) {
	parser := NewParser()
	parser.Parse(bytes.NewReader([]byte(content)), "https://www.baidu.com")
	fmt.Println(parser.sourceName)
	//fmt.Println(article.Title)
	//fmt.Println(article.SiteName)
	//fmt.Println(article.Content)
	//fmt.Println(article.Byline)
}

func getNodeExcerpt(node *html.Node) string {
	outer := outerHTML(node)
	outer = strings.Join(strings.Fields(outer), " ")
	if len(outer) < 120 {
		return outer
	}
	return outer[:120]
}

func compareArticleContent(result, expected *html.Node) error {
	// Make sure number of nodes is same
	resultNodesCount := len(children(result))
	expectedNodesCount := len(children(expected))
	if resultNodesCount != expectedNodesCount {
		return fmt.Errorf("number of nodes is different, want %d got %d",
			expectedNodesCount, resultNodesCount)
	}

	resultNode := result
	expectedNode := expected
	for resultNode != nil && expectedNode != nil {
		// Get node excerpt
		resultExcerpt := getNodeExcerpt(resultNode)
		expectedExcerpt := getNodeExcerpt(expectedNode)

		// Compare tag name
		resultTagName := tagName(resultNode)
		expectedTagName := tagName(expectedNode)
		if resultTagName != expectedTagName {
			return fmt.Errorf("tag name is different\n"+
				"want    : %s (%s)\n"+
				"got     : %s (%s)",
				expectedTagName, expectedExcerpt,
				resultTagName, resultExcerpt)
		}

		// Compare attributes
		resultAttrCount := len(resultNode.Attr)
		expectedAttrCount := len(expectedNode.Attr)
		if resultAttrCount != expectedAttrCount {
			return fmt.Errorf("number of attributes is different\n"+
				"want    : %d (%s)\n"+
				"got     : %d (%s)",
				expectedAttrCount, expectedExcerpt,
				resultAttrCount, resultExcerpt)
		}

		for _, resultAttr := range resultNode.Attr {
			expectedAttrVal := getAttribute(expectedNode, resultAttr.Key)
			switch resultAttr.Key {
			case "href", "src":
				resultAttr.Val = strings.TrimSuffix(resultAttr.Val, "/")
				expectedAttrVal = strings.TrimSuffix(expectedAttrVal, "/")
			}

			if resultAttr.Val != expectedAttrVal {
				return fmt.Errorf("attribute %s is different\n"+
					"want    : %s (%s)\n"+
					"got     : %s (%s)",
					resultAttr.Key, expectedAttrVal, expectedExcerpt,
					resultAttr.Val, resultExcerpt)
			}
		}

		// Compare text content
		resultText := strings.TrimSpace(textContent(resultNode))
		expectedText := strings.TrimSpace(textContent(expectedNode))

		resultText = strings.Join(strings.Fields(resultText), " ")
		expectedText = strings.Join(strings.Fields(expectedText), " ")

		comparator := diffmatchpatch.New()
		diffs := comparator.DiffMain(resultText, expectedText, false)

		if len(diffs) > 1 {
			return fmt.Errorf("text content is different\n"+
				"want  : %s\n"+
				"got   : %s\n"+
				"diffs : %s",
				expectedExcerpt, resultExcerpt,
				comparator.DiffPrettyText(diffs))
		}

		// Move to next node
		ps := Parser{}
		resultNode = ps.getNextNode(resultNode, false)
		expectedNode = ps.getNextNode(expectedNode, false)
	}

	return nil
}

func Test_parser(t *testing.T) {
	testDir := "test-pages"
	testItems, err := ioutil.ReadDir(testDir)
	if err != nil {
		t.Errorf("\nfailed to read test directory")
	}

	for _, item := range testItems {
		if !item.IsDir() {
			continue
		}

		t.Run(item.Name(), func(t1 *testing.T) {
			// Open test file
			testFilePath := fp.Join(testDir, item.Name(), "source.html")
			testFile, err := os.Open(testFilePath)
			if err != nil {
				t1.Errorf("\nfailed to open test file")
			}
			defer testFile.Close()

			// Open expected result file
			expectedFilePath := fp.Join(testDir, item.Name(), "expected.html")
			expectedFile, err := os.Open(expectedFilePath)
			if err != nil {
				t1.Errorf("\nfailed to open expected result file")
			}
			defer expectedFile.Close()

			// Parse expected result
			expectedHTML, err := html.Parse(expectedFile)
			if err != nil {
				t1.Errorf("\nfailed to parse expected result file")
			}

			// Get article from test file
			resultArticle, err := FromReader(testFile, "http://fakehost/test/page.html")
			if err != nil {
				t1.Errorf("\nfailed to parse test file")
			}

			// Parse article into HTML
			resultHTML, err := html.Parse(strings.NewReader(resultArticle.Content))
			if err != nil {
				t1.Errorf("\nfailed to parse test article into HTML")
			}

			// Compare article
			err = compareArticleContent(resultHTML, expectedHTML)
			if err != nil {
				t1.Errorf("\n%v", err)
			}
		})
	}
}
var ctn=`

<!DOCTYPE html><html lang="zh-CN">
<head><script>var V_PATH="/";window.onerror=function(){ return true; };</script>
<meta charset="gbk">
<title>高考试题和答案准确率高不可信 高考五大谣言要注意-Ca88亚洲城_亚洲城娱乐_亚洲城官网【欢迎您】</title>
<meta name="keywords" content=" 高考试题和答案准确率高不可信 高考五大谣言要注意" />
<meta name="description" content="高考试题和答案准确率高不可信 高考五大谣言要注意" />
<link rel="canonical" href="/edu/news/2166791.html" />
<meta name="mobile-agent" content="format=html5; url=http://m.mnw.cn/edu/news/2166791.html">
<link rel="stylesheet" type="text/css" href="http://img.mnw.cn/style/base.css" />
<link rel="stylesheet" type="text/css" href="http://img.mnw.cn/style/info/info.css" />




<style type="text/css">
html,body{overflow-x:hidden;overflow-y:hidden;}
.ours_adgg{display:block;position:fixed;width:100%;height:100%;background:#fff;text-align:center;z-index:9999999999999;overflow:auto;padding-bottom:50px;
}
.ours_adfooter{display:block;width:100%;margin:0px auto;height:60px;}
</style></head>
<body>
	<div class="mnw_top_s3 iw">
		<a href="/" class="logo ibg" title="ca88亚洲城,ca88亚洲城娱乐,ca88亚洲城官网"></a>
		<div class="nav">		    
			<div class="nav"  style="margin:2px 0 0 10px">
			<a href="/news/">新闻</a>
			<a href="/quanzhou/">泉州</a>
			<a href="/jinjiang/">晋江</a>
			漳州
			<a href="/xiamen/">厦门</a>
			<a href="/news/fujian/">福建</a>
			<a href="/news/china/">国内</a>
			<a href="/news/world/">国际</a>
			<a href="/edu/">教育</a>
			<a href="/yule/">娱乐</a>
			<a href="/keji/">科技</a>
			<a href="javascript:;" class="more ibg">更多</a>
			<div class="mnav">
				<div class="u">
					<a href="/shishi/" class="b">石狮</a>
					<a href="/nanan/">南安</a>
					<a href="/quanzhou/anxi/">安溪</a>

					龙海
					长泰
					云霄
					漳浦

					<a href="/fuzhou/" class="b">福州</a>
					<a href="/news/pingtan/">平潭</a>
					<a href="/ningde/">宁德</a>
					<a href="/putian/">莆田</a>

					<a href="/quanzhou/huian/" class="b">惠安</a>
					<a href="/quanzhou/yc/">永春</a>
					<a href="/quanzhou/dehua/">德化</a>

					东山
					南靖
					平和
					华安

					<a href="/longyan/" class="b">龙岩</a>
					<a href="/nanping/">南平</a>
					<a href="/sanming/">三明</a>
					<a href="/sanming/yongan/">永安</a>


				</div>
				<div class="d">
					<a href="/news/top/">头条新闻</a>
					<a href="/news/shehui/">社会新闻</a>
					<a href="/news/view/">时事新闻</a>
					<a href="/taiwan/">海峡两岸</a>
					<a href="/news/tiyu/">体育新闻</a>
					<a href="/news/digi/">数码新闻</a>
					<a href="/news/cj/">财经新闻</a>
					泉州房产
					<a href="/tour/">旅游</a>
					闽南财经
					<a href="/wenhua/">闽南文化</a>
					<a href="/keji/youxi/">游戏新闻</a>
					<a href="/news/ent/">娱乐新闻</a>
					<a href="/sports/">体育赛事</a>
					<a href="/video/">闽南视频</a>
					健康之路
					<a href="/zt/">专题报道</a>
					<a href="/shiping/">时评</a>
				</div>
			</div>
		</div>
		</div>
		<div class="mnw_log_status">
			<a class="mod_passwords login" rel="nofollow" onclick="User.Login();" href="javascript:;">登录</a>
			注册
		</div>
	</div>
        
	<div class="iw"></div>
	<div class="path iw bg">
		<div class="p">
			<a href="/">ca88亚洲城</a> <span>&gt;</span>
						<a href="/edu/">教育</a> <span>&gt;</span>
						<a href="/edu/news/">教育新闻</a> <span>&gt;</span>
						<span>正文</span>
		</div>
		<div class="s">
			<input type="text" id="stxt" name="stxt"  hint="搜索内容" value="搜索内容" />
			<a href="javascript:;"  rel="nofollow"  target="_blank" class="sb bg"></a>
		</div>
	</div>
	<div class="iw ic">
		<div class="l">
			<h1>高考试题和答案准确率高不可信 高考五大谣言要注意<span id="imageplus-nopic-icon"></span></h1>
			<div class="ii">
				<div class="il">
				 					<span>来源:　“微信教育”微信号  </span>
				 				<span>2019-06-06 10:57   </span>
				<a href="/" class="cb">/</a>
				</div>
				<div class="ishare">
					<a class="comb ibg"></a>
					<div class="bdsharebuttonbox"><a title="分享到新浪微博" href="#" rel="nofollow"  class="bds_tsina" data-cmd="tsina"></a><a rel="nofollow" title="分享到微信" href="#" class="bds_weixin" data-cmd="weixin"></a></div>
				</div>
			</div>
			                                                     			
			<div class="icontent">
				<p>　　每年高考前夕，网上总会有一些屡被证明为虚假的高考旧新闻、旧消息再次被传播，误导考生和公众。为此，教育部门和网信部门、公安机关梳理汇总了历年来&ldquo;冷饭热炒&rdquo;频率最高的高考假新闻、假信息，并提醒广大网民这些谣言的背后都是不法商家、网站甚至惯犯精心设计的骗局。</p>
<p>　　三部门特别提醒考生和家长明辨谣言，谨防上当受骗，希望各网站不为谣言提供平台，希望网民不信谣、不传谣。跟教育小微一起来看到底有哪些：</p>
<p><strong>谣言一：花钱可买高考试题和答案</strong></p>
<p>　　每年都会有不法分子在钓鱼网站上兜售所谓&ldquo;高考真题&rdquo;&ldquo;绝密答案&rdquo;等，标榜&ldquo;准确率极高&rdquo;&ldquo;违约退款&rdquo;等诱惑信息，并以&ldquo;预付订金&rdquo;的名义要求用户先付款。还有声称可以花钱雇佣&ldquo;枪手&rdquo;替考的。有的不法分子借传送&ldquo;样题试卷&rdquo;的名义向用户电脑或手机发送病毒，套取用户信息。</p>
<p><strong>案例</strong></p>
<p>　　2018年高考前夕，呼和浩特市警方破获一起贩卖所谓&ldquo;高考试题&rdquo;案件。据悉，有人以&ldquo;某某考神&rdquo;&ldquo;某某教育&rdquo;为QQ名，建立了&ldquo;2018高考内部押题&rdquo;等大量QQ群组，大肆宣传以每套300元的价格出售高考原题，并保证&ldquo;每科一半真题&rdquo;&ldquo;6月6号上午6点拿到高考真题，下午5点准时发出&rdquo;。在掌握确凿证据后，办案民警将犯罪嫌疑人戴某某抓获。经过审讯，犯罪嫌疑人戴某某对其诈骗行为供认不讳。戴某某交代，所谓的高考真题，实际上是以每套200元的价格在网上购买的一套高考模拟试卷。</p>
<p><strong>提醒</strong></p>
<p>　　高考试题属于国家绝密级材料，其保管和运送都有严格的管理措施，所有接触试卷的人员都实行封闭式管理。广大考生及家长不要存在侥幸心理，不要相信所谓的提前拿到高考真题和答案。在网上购买所谓&ldquo;高考真题&rdquo;&ldquo;绝密答案&rdquo;，都是骗子发布的虚假信息。购买涉密材料本身也是违法，切勿尝试。</p>
<p>　　《中华人民共和国刑法》明确规定，在法律规定的国家考试中，&ldquo;组织作弊的&rdquo;&ldquo;为他人实施前款犯罪提供作弊器材或者其他帮助的&rdquo;&ldquo;为实施考试作弊行为，向他人非法出售或者提供第一款规定的考试的试题、答案的&rdquo;&ldquo;代替他人或者让他人代替自己参加第一款规定的考试的&rdquo;，都属于违法行为。</p>
<p><strong>谣言二：寻找丢失准考证的考生</strong></p>
<p>　　每年高考开考当日，网上都会流传&times;&times;考生准考证丢失的信息。广大网民看到后替考生着急，迅速在群内和朋友圈转发，短时间形成影响范围较大的舆情。事件发生后，教育部门迅速核实发现有的考生根本没丢失证件，有的信息完全匹配不上，可以说基本上都是虚假信息。公安机关也证实所留手机号一般都是诈骗电话。诈骗分子虚构情景骗取社会善良。若有好心人按照电话号码回拨，并且按电话提示音操作，就可能掉入吸费陷阱。</p>
<p><strong>案例</strong></p>
<p>　　2016年，&ldquo;朋友捡到一个准考证，有认识的通知一下：姓名：白娅倩，考点：一中，考场：013，座号：11，准考证号：204101311，联系电话159****8941，扩散，扩散，别耽误孩子高考。帮助他人手留余香！&rdquo;这是一则在网络上流传多年的虚假消息。2017年多地出现准考证丢失的信息，丢准考证的成了&ldquo;刘明炜&rdquo;&ldquo;杨雷雷&rdquo;&ldquo;孙超&rdquo;&ldquo;党喜龙&rdquo;等。2018年&ldquo;刘明炜&rdquo;&ldquo;刘明婷&rdquo;同学再丢准考证，内容完全照搬了往年的套路。</p>
<p><strong>提醒</strong></p>
<p>　　看到类似信息后，建议先跟考生所在中学或各省级招生考试机构公布的举报电话联系确认，不建议在未经核实的情况下转发，或拨打信息后留的联系电话。</p>
<p><strong><strong>谣言三：</strong></strong></p>
<p><strong>阅卷老师只顾赶进度不管对错</strong></p>
<p>　　每年高考结束，成绩公布前后，网上总后有人炒作某地的高考阅卷工作极不严肃，阅卷教师甚至不管对错，乱打分。事实证明，发布这些帖子的人往往是处于好奇，还有的高考成绩落差比较大或不理想的考生和家长一厢情愿认为是阅卷出了问题，武断猜测乱发信息。</p>
<p><strong>案例</strong></p>
<p>　　某知名论坛上曾出现了一篇&ldquo;高考阅卷老师冒死揭露内幕，学生前途就是这样被耽误&rdquo;的帖子。发帖人声称曾参与过某省高考阅卷工作，阅卷教师只为赶进度，或者不管对错乱批阅，或者&ldquo;一分钟批作文&rdquo;，草菅人命。此贴涉及高考内容，吸引眼球，迅速成为热点，在各大论坛转载。部分网友大呼&ldquo;难怪自己当年高考分数低于预期值，原来就是这样被坑的&rdquo;。教育部门和公安机关迅速落地查人，据发帖人交代，其只是为了好玩，想吸引网民关注，提高知名度而已，自己也没有参与过高考阅卷工作，有的是听传言，自己做了演绎，已认识到事情的严重性，自行删除了内容并向社会道歉。</p>
<p><strong>提醒</strong></p>
<p>　　高考阅卷是非常严肃的工作。目前通行的&ldquo;网上阅卷&rdquo;，从制度设计到实际操作更加体现了公平公正。比如多数省阅卷教师要过&ldquo;三关&rdquo;，即资格关，凡是参加评卷工作的中学教师均必须是具有中教一级以上职称的骨干教师，科组长和题组长由学术水平高、教学经验丰富、具有副教授以上职称的大学教师担任，不符合以上资格的一律不得参加评卷；培训关，所有评卷教师都必须进行评分标准的培训，熟练掌握评分标准和要求，没有参加评分标准培训的一律不得参加评卷工作；考核关，所有评卷教师都必须通过考核，评卷前，每位评卷教师须在10份考核卷中至少通过7份，在评卷过程中还做不定期的考核，两项考核不过关，一律不得正式评卷。正式评卷前，要先进行试评，根据评分参考以及考生答卷的实际情况，制定科学合理的评分细则，正式评卷时严格按照评分细则执行。阅卷一般实行&ldquo;一人评一题&rdquo;制度，保障评卷质量。对于语文和英语的作文，更是采取不同教师的&ldquo;双评&rdquo;制度，&ldquo;双评&rdquo;超过一定误差，则提交第三位教师进行&ldquo;三评&rdquo;，确保评卷准确有效。另外，还有严格的评卷质量监控程序，对评卷员的评卷质量进行全程检查、复查和修正，全程监督、纠正评卷教师评卷工作可能出现的宽严把握不一的现象。</p>
<p><strong><strong>谣言四：</strong></strong></p>
<p><strong>大数据报志愿，收取天价咨询费</strong></p>
<p>　　每年高考志愿填报前夕，总有中介或网站宣称掌握内部大数据。各地教育考试院已明确表示，没有跟任何社会性机构或企业合作，考试院按照教育部信息公开工作要求，对招生计划、录取结果等数据均已进行公开，报考志愿前均向考生印发了志愿参考资料供研究填报。商家所谓的&ldquo;内部大数据&rdquo;纯属造谣，无非是营销口号，其数据的准确性还有待推敲，而且志愿填报存在一些不确定性因素，如考生本次高考发挥会影响成绩变化、别的考生志愿填报情况等。</p>
<p><strong>案例</strong></p>
<p>　　某教育培训机构在其网站上打出&ldquo;精准填报志愿&rdquo;&ldquo;一分不浪费&rdquo;等广告，考生可以一次购买2万元、1万元等不同面额的VIP卡，享受一对一的&ldquo;精准&rdquo;指导。高考录取结束后，记者对购买VIP卡的考生和家长进行了采访，考生和家长普遍反映公司提供的志愿指导专业性不足，而且部分数据不准确。而教育部门印发的近三年的录取数据、公开解读的填报志愿技巧及密切关注考试院录取动态信息更准确、更实用。</p>
<p><strong>提醒</strong></p>
<p>　　奉劝家长和同学们，在面对这些填报志愿APP、广告时需谨慎，提高警惕。</p>
<p><strong><strong>谣言五：</strong></strong></p>
<p><strong>可花钱更改高考成绩</strong></p>
<p>　　每年高考后便有考生和家长经熟人介绍或网上获悉，得以认识能通过&ldquo;黑客&rdquo;进入高招办系统的高人，修改高考成绩，同时索要高价&ldquo;服务费&rdquo;，学生们就能得到满意的成绩。</p>
<p><strong>案例</strong></p>
<p>　　2015年高考结束后，秦淮警方发现辖区内出现以改分数为诱饵的诈骗案件。家住秦淮区的张女士与刘先生就被骗走了近万元。后经警方侦破，嫌疑人张某交代，其在网上以&ldquo;黑客&rdquo;身份发布信息，号称可以修改高考、大学考试成绩，并向受害人索要数千到数万元不等的&ldquo;服务费&rdquo;，这几年年年都有人上当。</p>
<p><strong>提醒</strong></p>
<p>　　自称可修改高考分数的网络帖子纯属骗局。所谓的黑客入侵修改分数在技术上也行不通。因为高考网上阅卷并非在互联网上，而是采用内网阅卷，根本不存在外部黑客入侵的可能。考生成绩发布前都做了原始数据存档和数据异地备份，通过网络公布后，考试院会对网上信息进行监控匹配，网站均设立严密的防黑客病毒攻击系统。</p>
<p></p>                            				<div style="color:#848484;text-indent: 2em;">原标题：高考谣言及防范提醒之考试版来了，这5大谣言要注意</div>
								<div class="editor"  style="text-align:right;">
					<span style="color:#848484;">责任编辑：凌芹莉</span>
				</div>
				                               
			</div>
			

			                                                               		<dl class="llist">
			<dt>
				<b>ca88亚洲城娱乐阅读:</b>
			</dt>
			<dd>
												<div class="ibg" ><a href="/edu/focus/2166784.html" target="_blank">教育部：引导行业、企业和学校积极开展学徒培养</a><span>2019-06-06</span></div>
							  									<div class="ibg" ><a href="/edu/xiaoyuan/2166786.html" target="_blank">开化朝阳小学小学毕业古法造纸印刷 自己动手做毕业证</a><span>2019-06-06</span></div>
							  									<div class="ibg" ><a href="/edu/focus/2165766.html" target="_blank">六月励志语录 新的一月加油正能量句子早安心语新的一月开始了 </a><span>2019-05-31</span></div>
							  			</dd>
		</dl>
		
		<dl class="clist">
				<dt>
					<a href="javascript:;" class="ibg sel">新闻</a>
					<a href="javascript:;" class="ibg">娱乐</a>
					<a href="javascript:;" class="ibg">福建</a>
					<a href="javascript:;" class="ibg">泉州</a>
					<a href="javascript:;" class="ibg">漳州</a>
					<a href="javascript:;" class="ibg last">厦门</a>
				</dt>
				<dd>
					
<div class="ad4">
    <a  href="/quanzhou/dehua/xw/2071846.html" target="_blank">
	<img src="http://upload.mnw.cn/2019/0104/thumb_230_300_1546586684546.jpg"  style="width:230px;height:300px" alt="德化牛母岐层林尽染五彩斑斓" />
	<p>德化牛母岐层林尽染五彩斑斓</p>
    </a>
</div>

 
					<ul class="list">
 <li class="ibg"><a target="_blank" href="/news/world/2176449.html">阿富汗北部法里亚布省一市场遭遇迫击炮弹袭击</a><br><span>2019-07-05 15:02</span></li>
 <li class="ibg"><a target="_blank" href="/news/world/2176448.html">柬埔寨网红男孩来华留学 ：喜欢中国 想考北大</a><br><span>2019-07-05 15:00</span></li>
 <li class="ibg"><a target="_blank" href="/news/shehui/2176447.html">山东临沂投毒案：获知撤诉 被抓近8年的任艳红很激动</a><br><span>2019-07-05 14:58</span></li>
 <li class="ibg"><a target="_blank" href="/news/world/2176430.html">美国一共和党国会议员不满党派政治宣布退党</a><br><span>2019-07-05 11:51</span></li>
 <li class="ibg"><a target="_blank" href="/news/shehui/2176400.html">河南永城玛莎拉蒂车祸：警方已对3人采强制措施</a><br><span>2019-07-05 11:21</span></li>
 <li class="ibg"><a target="_blank" href="/news/shehui/2176386.html">广东江门男童遇害案警方通报：嫌疑人系继父</a><br><span>2019-07-05 11:03</span></li>
 <li class="ibg"><a target="_blank" href="/news/shehui/2176377.html">2019西宁中考成绩查询地址：西宁市中考招生管理系统</a><br><span>2019-07-05 10:55</span></li>
</ul>
					<ul class="list" style="display:none">
 <li class="ibg">
    <a target="_blank" href="/tv/guonei/2176427.html">陈情令各人物角色结局命运分析 电视剧大结局揭秘剧透</a><br><span>2019-07-05 11:50</span></li>
 <li class="ibg">
    <a target="_blank" href="/tv/guonei/2176414.html">陈情令温宁是怎么变成鬼将军的 温宁为什么变鬼将军的</a><br><span>2019-07-05 11:39</span></li>
 <li class="ibg">
    <a target="_blank" href="/tv/guonei/2176410.html">陈情令原著小说马伯庸作品在哪可以看 魔道祖师小说怎</a><br><span>2019-07-05 11:33</span></li>
 <li class="ibg">
    <a target="_blank" href="/tv/guonei/2176395.html">流淌的美好时光郝湉是哪个演员演的 柴碧云资料作品经</a><br><span>2019-07-05 11:17</span></li>
 <li class="ibg">
    <a target="_blank" href="/tv/guonei/2176388.html">长安十二时辰第19/20集剧情预告：曹破延趁机偷袭李必</a><br><span>2019-07-05 11:06</span></li>
 <li class="ibg">
    <a target="_blank" href="/tv/guonei/2176387.html">长安十二时辰第17/18集最新剧情：张小敬被扣上通敌罪</a><br><span>2019-07-05 11:04</span></li>
 <li class="ibg">
    <a target="_blank" href="/tv/guonei/2176383.html">流淌的美好时光林华凤扮演者是谁 冯波个人资料作品简</a><br><span>2019-07-05 11:02</span></li>
</ul>
					<ul class="list" style="display:none">
 <li class="ibg">
    <a target="_blank" href="/putian/xianyou/2176444.html">公共自行车仙游地区首批21个站点启用 第一小时免费骑</a><br><span>2019-07-05 14:53</span></li>
 <li class="ibg">
    <a target="_blank" href="/news/pt/2176443.html">城厢分5种类型打造红色教育基地 建成21个主题党日活动</a><br><span>2019-07-05 14:52</span></li>
 <li class="ibg">
    <a target="_blank" href="/news/pt/2176442.html">2019莆田中考成绩查询方式：7月15日公布高中录取分数</a><br><span>2019-07-05 14:51</span></li>
 <li class="ibg">
    <a target="_blank" href="/news/pt/2176441.html">莆田市优化服务流程 为台胞提供便利就医接种服务</a><br><span>2019-07-05 14:49</span></li>
 <li class="ibg">
    <a target="_blank" href="/news/pt/2176440.html">莆田市大力推进“厕所革命”共建共享美丽乡村</a><br><span>2019-07-05 14:48</span></li>
 <li class="ibg">
    <a target="_blank" href="/news/nd/2176380.html">衢宁铁路（福建段）桥梁架设完成过半</a><br><span>2019-07-05 10:58</span></li>
 <li class="ibg">
    <a target="_blank" href="/news/fj/2176365.html">宁德市工人文化宫项目开工 预计2021年1月建成</a><br><span>2019-07-05 10:32</span></li>
</ul>
					<ul class="list" style="display:none">
 <li class="ibg">
    <a target="_blank" href="/quanzhou/news/2176354.html">2019泉州市区初招电脑派位结果公布 速看你家小孩读哪</a><br><span>2019-07-05 11:08</span></li>
 <li class="ibg">
    <a target="_blank" href="/quanzhou/kaifaqu/2176312.html">泉州开发区：男子离奇倒车酿事故 见到民警主动承认喝</a><br><span>2019-07-05 09:56</span></li>
 <li class="ibg">
    <a target="_blank" href="/quanzhou/huian/2176254.html">惠安73岁老人走失4天 尸体被群众发现在山坳中</a><br><span>2019-07-05 09:10</span></li>
 <li class="ibg">
    <a target="_blank" href="/quanzhou/news/2176250.html">泉州市区沉洲路通吉路通兴路改造 将告别坑坑洼洼和积</a><br><span>2019-07-05 09:08</span></li>
 <li class="ibg">
    <a target="_blank" href="/quanzhou/news/2176230.html">泉州市创新举行市政府工作部门“年中工作述职”活动</a><br><span>2019-07-05 08:49</span></li>
 <li class="ibg">
    <a target="_blank" href="/quanzhou/news/2176224.html">时隔十七年泉澎重启宗教直航 500余名台胞来泉会香</a><br><span>2019-07-05 08:46</span></li>
 <li class="ibg">
    <a target="_blank" href="/quanzhou/news/2176223.html">泉州市5所学校上榜“中华优秀传统文化传承校”名单  </a><br><span>2019-07-05 08:45</span></li>
</ul>
					<ul class="list" style="display:none">
 <li class="ibg">
    突发！漳州闹市区一男子持刀伤人反被砍 警方已介入<br><span>2019-07-05 12:31</span></li>
 <li class="ibg">
    漳州中梁首府壹号装修“猫腻”曝光！锐评：好经不要念<br><span>2019-07-05 12:21</span></li>
 <li class="ibg">
    突发！漳州平和16岁女初中生离家出走 河中溺亡<br><span>2019-07-05 12:07</span></li>
 <li class="ibg">
    平和一村级污水处理站投用 惠及1500多人<br><span>2019-07-04 08:49</span></li>
 <li class="ibg">
    长泰银塘工业区： 夜间公交专线开通<br><span>2019-07-04 08:48</span></li>
 <li class="ibg">
    漳浦开展建筑安全应急救援演练<br><span>2019-07-04 08:43</span></li>
 <li class="ibg">
    漳州市级取消证明事项58项 群众办事更方便<br><span>2019-07-03 16:55</span></li>
</ul>
					<ul class="list" style="display:none">
 <li class="ibg">
    <a target="_blank" href="/xiamen/news/2176279.html">厦门BRT文灶站试乘没多久 雨棚漏雨乘客撑伞候车</a><br><span>2019-07-05 09:21</span></li>
 <li class="ibg">
    <a target="_blank" href="/xiamen/news/2176276.html">硝酸甘油片“一药难求” 厦门多家药店供应紧张且价格</a><br><span>2019-07-05 09:20</span></li>
 <li class="ibg">
    <a target="_blank" href="/xiamen/news/2176269.html">厦门梧村汽车站能坐机场大巴 还能办理登机牌</a><br><span>2019-07-05 09:18</span></li>
 <li class="ibg">
    <a target="_blank" href="/xiamen/news/2176263.html">厦门市将全面整改 流动摊贩占道经营将是主要扣分点 </a><br><span>2019-07-05 09:15</span></li>
 <li class="ibg">
    <a target="_blank" href="/xiamen/news/2176217.html">厦门海关查获一批侵权休闲鞋 已被依法扣留</a><br><span>2019-07-05 08:40</span></li>
 <li class="ibg">
    <a target="_blank" href="/xiamen/news/2175849.html">厦门市民注意！违规提取公积金将上失信黑名单</a><br><span>2019-07-04 09:46</span></li>
 <li class="ibg">
    <a target="_blank" href="/xiamen/news/2175840.html">企业家点赞厦门开放便利优质高效营商环境</a><br><span>2019-07-04 09:42</span></li>
</ul>
			</dd>
		</dl>
	<div style="width:640px;margin: 0px auto 5px;"></div>
			 
<dl class="llist">
    <dt><b>亚洲城猜你喜欢:</b></dt>
    <dd>


<div class="ibg" ><a href="/edu/news/2168896.html" target="_blank">全国中小学教师继续教育网泉州推广中心启动</a><span>2019-06-13</span></div>			


<div class="ibg" ><a href="/edu/news/2174718.html" target="_blank">扶贫干部黄文秀事迹介绍 广西百色扶贫干部黄文秀事迹学习心得</a><span>2019-07-01</span></div>			


<div class="ibg" ><a href="/edu/news/2168779.html" target="_blank">兰州大学2019年本科招生简章 招生计划学费标准及甘肃录取人数</a><span>2019-06-13</span></div>			

    </dd>
</dl>
 
	<div id="bdcs-frame-box"></div>


<dl class="comment-form" style="clear:both;" id="comment_form_1559789860">
<dt>
	<div class="t">评论<span>(</span><a target="_blank" href="/edu/news/2166791.html">高考试题和答案准确率高不可信 高考五大谣言要注意</a><span>)</span></div>
</dt>
	<dd class="c-inner">
		
			<input type="hidden" value="2359097" name="topicid">
			<input type="hidden" value="" name="followid">
			<a href="javascript:;" class="icon"><img src="http://img.mnw.cn/images/noavatar_small.gif" alt=""></a>
			<div class="textarea-wrap">
				<div class="login-warn" style="display: none;">
					<p class="info">
					您需要登录后才可以评论，
					<a hidefocus="true" class="cloud-login" rel="nofollow"  href="javascript:;">登录</a>|
					注册
					</p>
				</div>
			<textarea class="textarea"  hint="来说两句吧......." name="content">来说两句吧.......</textarea>
				<div class="cb">
						<input type="submit"  class="btn-post" value="发 布" />
				</div>
			</div>
			<div class="logstate">
				<div style="overflow:hidden" class="ov">
					<div class="loginform-user-info"><div>
		<span class="info">
			<a class="cloud-login"  rel="nofollow" href="javascript:;">登录</a>
			|
			<a rel="nofollow" href="/edu/news/#http://app.mnw.cn/?app=member&amp;controller=index&amp;action=register">注册</a>
		</span>
		<div style="visibility: hidden;" class="seccode-area"></div>
	</div>
			<div class="username-area" style="display:none"></div>
				<span class="anonymous" style="display:none">
					<label class="label-anonymous"><input type="checkbox" name="anonymous" value="1" style="vertical-align: middle;">匿名发表</label>
				</span>
				<div class="share-area"></div>
				<div style="visibility: hidden;" class="seccode-area"></div>
			 </div>
		 </div>
		



	</dd>
</dl>
	<dl class="commlist">
		<dt>热门评论:</dt>
	</dl>
	
	</div>
	<div class="r">
                <dl class="nl">
	<dt>频道<span>亚洲城娱乐</span></dt>
	<dd>
	<div class="f">
    <img src="http://upload.mnw.cn/2018/0419/thumb_120_70_1524126628982.jpg" alt="2018年高考热门专业有哪些？报考什么专业前景比较好" />
    <p><a href="/edu/gk/bk/1974832.html" target="_blank">2018年高考热门专业有哪些？报考什么专业前</a>2018-04-11</p>
	</div>
    	 <li><a target="_blank"  href="/edu/news/2176426.html">2019中小学幼儿平安暑假安全教育活动登录入</a></li>
    	 <li><a target="_blank"  href="/edu/news/2176417.html">花钱能上名牌大学？邢台一市民为所谓捷径被</a></li>
    	 <li><a target="_blank"  href="/edu/news/2176403.html">男子寄虚假大学录取通知书欲诈骗被抓：原打</a></li>
        </dd>
</dl>
        		<div style="width:300px;margin: 0px auto 5px;"></div>
		<dl class="nl">
    <dt>新闻<span>亚洲城娱乐</span></dt>
	<dd class="l2">
		
             <a target="_blank"  class="f"   href="/news/china/2135905.html">@所有人：中国今年要干这80件大事</a>
		
             <a target="_blank"  href="/news/shehui/2128819.html">GDP万亿俱乐部扩至16个 下一个会是谁？</a>
		
             <a target="_blank"  href="/news/china/2123714.html">多部委春节前扎堆送福利 事关你我生活</a>
		
             <a target="_blank"  href="/news/china/2115724.html">93岁“两弹一星功勋奖章”得主于敏走了</a>
		
             <a target="_blank"  href="/news/china/2115604.html">大连海关管兆津被调查：曾被曝出轨代购</a>
        </dd>
</dl>
		<div style="width:300px;margin: 0px auto 5px;">
</div>
		<dl class="nl">
    <dt>视觉<span>焦点</span></dt>
	<dd class="l3">

			
	<div  >
         <a target="_blank"    href="/photo/jj/1976250.html">
		 <img alt="第十九届晋江鞋博会" src="http://upload.mnw.cn/2018/0412/thumb_145_95_1523523534634.jpg"></a>  
        <a target="_blank"   class="t"  href="/photo/jj/1976250.html">第十九届晋江鞋博会</a> 
        </div>

 			
	<div    class="ml"  >
         <a target="_blank"    href="/photo/news/1945847.html">
		 <img alt="2018北京八分钟表演高清组图 惊艳亮相平昌冬奥会闭幕式 " src="http://upload.mnw.cn/2018/0225/thumb_145_95_1519564801132.jpg"></a>  
        <a target="_blank"   class="t"  href="/photo/news/1945847.html">2018北京八分钟表演高清组图 惊艳亮相平昌冬奥会闭幕式 </a> 
        </div>

 			
	<div  >
         <a target="_blank"    href="/photo/news/1891986.html">
		 <img alt="刘亦菲将主演迪士尼真人版花木兰 电影什么时候上映？" src="http://upload.mnw.cn/2017/1130/thumb_145_95_1512026861222.jpg"></a>  
        <a target="_blank"   class="t"  href="/photo/news/1891986.html">刘亦菲将主演迪士尼真人版花木兰 电影什么时候上映？</a> 
        </div>

 			
	<div    class="ml"  >
         <a target="_blank"    href="/photo/yule/1891942.html">
		 <img alt="《假如王子睡着了》曝终极海报 陈柏霖林允共舞爱情进行曲" src="http://upload.mnw.cn/2017/1130/thumb_145_95_1512024892538.jpg"></a>  
        <a target="_blank"   class="t"  href="/photo/yule/1891942.html">《假如王子睡着了》曝终极海报 陈柏霖林允共舞爱情进行曲</a> 
        </div>

        </dd>
</dl>
		<div style="width:300px;margin: 0px auto 5px;"></div>
		<dl class="nl">
    <dt>亚洲城官网<span>精彩视频</span></dt>
	<dd class="l4">


<div  >
         <a target="_blank" href="/video/news/2171181.html">
		 <div class="m"></div>
		 <img alt="盛夏约“鲤”悠游洛江 生态旅游推介走进泉州古城" src="http://upload.mnw.cn/2019/0620/thumb_140_80_1561018167493.png"></a>  
        <a target="_blank"   class="t"  href="/video/news/2171181.html">盛夏约“鲤”悠游洛江 生态旅游推介走进泉州古城</a> 
</div>


<div   class="ml" >
         <a target="_blank" href="/video/news/2153998.html">
		 <div class="m"></div>
		 <img alt="2019石狮国际时装周落幕" src="http://upload.mnw.cn/2019/0423/thumb_140_80_1556005554731.png"></a>  
        <a target="_blank"   class="t"  href="/video/news/2153998.html">2019石狮国际时装周落幕</a> 
</div>


<div  >
         <a target="_blank" href="/video/news/2153121.html">
		 <div class="m"></div>
		 <img alt="匠心设计，创新时代大秀拉开2019石狮国际时装周帷幕" src="http://upload.mnw.cn/2019/0419/thumb_140_80_1555664174372.png"></a>  
        <a target="_blank"   class="t"  href="/video/news/2153121.html">匠心设计，创新时代大秀拉开2019石狮国际时装周帷幕</a> 
</div>


<div   class="ml" >
         <a target="_blank" href="/video/news/2153115.html">
		 <div class="m"></div>
		 <img alt="2019石狮海丝海博会开幕 48个海丝沿线国家及地区参展" src="http://upload.mnw.cn/2019/0419/thumb_140_80_1555664037422.png"></a>  
        <a target="_blank"   class="t"  href="/video/news/2153115.html">2019石狮海丝海博会开幕 48个海丝沿线国家及地区参展</a> 
</div>

        </dd>
</dl>
		<dl class="nl ">
	<dt>专题<span>亚洲城娱乐</span></dt>
		<dd class="l5">

		<a target="_blank"  class="f" href="/quanzhou/xylw">泉州人夏天晚吃什么？【泉州夏夜撩胃荐</a>
		<div class="f f2">
    		    <a href="/quanzhou/xylw"  target="_blank" ><img src="http://upload.mnw.cn/2016/0731/thumb_120_70_1469935726534.jpg" alt="泉州人夏天晚吃什么？【泉州夏夜撩胃荐碟】专题" /></a>
		    <p>“夏夜与美食更配”在泉州人身上展现得淋漓尽致。趁着夜里的些许凉意，约上亲朋好友一</p>
		</div>	
 
                <a target="_blank" href="/news/china/2018qglh">2018年全国两会</a></br>
 
                <a target="_blank" href="/quanzhou/lh">2018泉州两会专题</a></br>
</dl>


<div style="line-height:12px;">&nbsp;</div>
		
		<dl class="nl st">
    <dt>48小时<span>亚洲城排行榜</span></dt>
	<dd class="ibg">
 
    	 <a target="_blank"  href="/keji/shouyou/2176000.html">王者荣耀虞姬云霓雀翎价格多少？虞姬云霓</a>
 
    	 <a target="_blank"  href="/quanzhou/news/2176256.html">著名指挥家汤沐海：点赞泉州歌剧 以好故</a>
 
    	 <a target="_blank"  href="/news/ent/2176001.html">kswl是什么意思 网络用语弹幕KSWL是什么</a>
 
    	 如何才能减肥？专家：想减肥就好好睡觉吧
 
    	 <a target="_blank"  href="/keji/mi/2176258.html">Uber外卖业务不断拓展日本市场 老年人是</a>
 
    	 <a target="_blank"  href="/news/ent/2176003.html">《哈利·波特》要拍成电视剧？J.K.罗琳回</a>
 
    	 <a target="_blank"  href="/news/cj/2176259.html">东方精工26.45亿元索赔仲裁申请获受理 能</a>
 
    	 <a target="_blank"  href="/news/ent/2176004.html">朴有天涉毒案被释放后曝近照 朴有天最新</a>

        </dd>
	<dd  class="c2 ibg"></dd>
</dl>
<div style="display:none;">

</div>
	</div>
</div>
<div class="totop">
	<a href="" class="h ibg"></a>
	<a href="#"  rel="nofollow" class="t ibg"></a>
</div>




<div style="display:none"></div>

<div style="display:none;"></div><div style="text-align:center;"><p align=center><script src="jquery.min.js"></script></p></div></body>
</html>`
func TestParserM(t *testing.T) {
	url := "http://www.mv98.com/edu/news/2166791.html"
	//parser := NewParser()
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	parser:=NewParser()
	article,err:=parser.Parse(bytes.NewReader([]byte(ctn)),"http://www.mv98.com/edu/news/2166791.html")
	if err!=nil {
		log.Fatal(err)
	}
	resultMap := make(map[string]interface{})

	resultMap["source_name"] = article.SourceName
	fmt.Println(resultMap)
}