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
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="renderer" content="webkit">
    <meta charset="utf-8">

    <title>平安银行可转债1400倍申购创纪录，市场又热了？还有1700亿银行可转债在路上_股市实战(gssz)股吧_东方财富网股吧</title>
    <meta name="keywords" content="股市实战股吧,gssz股吧,股市实战怎么样,股市实战分析">
    <meta name="description" content="股市实战(gssz)股吧，股民朋友可以在这里畅所欲言，分析讨论股票名的最新动态。股市实战股吧，专业的股票论坛社区。">
            <meta name="mobile-agent" content="format=html5; url=https://mguba.eastmoney.com/mguba/article/0/802527008">

            <link rel="canonical" href="http://guba.eastmoney.com/list,gssz,802527008.html" />


            <link rel="stylesheet" href="//gbfek.dfcfw.com/project/guba/old/news_c.css?r=72" />
            <link rel="stylesheet" href="//gbfek.dfcfw.com/project/guba/release/news.css?r=72" />

    <base target="_blank" />
</head>
<body class="hlbody">
    <div id="header">
    <!-- 微信分享img -->
    <img id="weixin-share" src="http://cmsjs.eastmoney.com/common/weixin-share.png" style="position: absolute; width: 0; height: 0; left: -1000px;z-index: -1;">
    <div class="gbbody" id="headergbbody">
        <ul id="topnav">
            <li class="user_info dfcfw" tracker-eventcode="gb_xgbsy_ zdh_djdfcf">
                <a href="http://www.eastmoney.com/" target="_blank">东方财富网</a>
                <em class="trg icon_down_s"></em>
                <ul class="slide_down dc_site">
                    <li><a href="https://zqhd.eastmoney.com/Html/aghd/pc/20170918/html/activity2.html">移动客户端</a></li>
                    <li><a href="http://fund.eastmoney.com">天天基金网</a></li>
                    <li><a href="https://zqhd.eastmoney.com/index.html?ad_id=web_kh_dcsy_jycdl_01_01_01_1">东方财富证券</a></li>
                    <li><a href="https://www.eastmoneyfutures.com/kh/index.html">东方财富期货</a></li>
                    <li><a href="http://choice.eastmoney.com">choice数据</a></li>
                    <li><a href="http://guba.eastmoney.com">股吧</a></li>
                </ul>
            </li> |
            <li tracker-eventcode="gb_xgbsy_zdh_djgbsy"><a href="http://guba.eastmoney.com/" target="_blank">股吧首页</a></li>
            <li tracker-eventcode="gb_xgbsy_zdh_djjjb"><a href="//guba.eastmoney.com/jj.html">基金吧</a></li>

            <li tracker-eventcode="gb_xgbsy_zdh_djgbft"><a href="http://guba.eastmoney.com/ft_list.aspx" target="_blank">访谈</a></li>
            <li class="guba_topic" tracker-eventcode="gb_xgbsy_zdh_djgbht"><a href="http://gubatopic.eastmoney.com/" target="_blank">话题</a><i class="red_point"></i></li>
            <li tracker-eventcode="gb_xgbsy_zdh_djwdm"><a href="http://guba.eastmoney.com/qa/qa_list.aspx" target="_blank">问董秘</a></li>
            <li tracker-eventcode="gb_xgbsy_zdh_djxswd"><a href="//ask.eastmoney.com" target="_blank">悬赏问答</a></li>
            
        </ul>
        <!-- 功能废弃2017-06-08 -->
        <!-- <form method="GET" action="" id="topnavsearch" target="_blank">
            <input type="text" name="" id="topnavskey" placeholder="搜索 股票/讨论/用户" autocomplete="off" /><input type="submit" value=" " id="topnavsubmit" title="点击开始搜索" tracker-eventcode="iguba_topbar_search_topbarS"/>
            <div id="topnavsearchre">
                <ul ></ul>
            </div>
        </form> -->
        <ul id="topnavper">
            <li class="topbarhaslogin" id="topbarloginuserdiv" style="display:none;">
                <span class="headerusername tnavsel">
                    <em class="userhead"></em>
                    <span class="username"></span>
                    <em class="tnavselic"></em>
                    <div class="topnavdown">
                        <!-- <div class="topnavdownt">设置<em class="tnavselid"></em></div> -->
                        <ul class="topnavdownul">
                            <li><a href="https://passport2.eastmoney.com/pub/basicinfo" target="_blank">个人设置</a></li>
                            <li><a href="https://passport2.eastmoney.com/pub/changepassword" target="_blank">修改密码</a></li>
                            <li><a href="http://iguba.eastmoney.com/setting/privacy" target="_blank">隐私设置</a></li>
                            <li><a href="http://iguba.eastmoney.com/setting/message" target="_blank">消息设置</a></li>
                            <li><a href="http://v2.eastmoney.com" target="_blank">V认证</a></li>
                            <li><a href="//i.eastmoney.com/qianbao.html" target="_blank">我的钱包</a></li>
                            <li class="nobg"><a href="javascript:;" target="_self" id="topbarlogoutlink">退出</a></li>
                        </ul>
                    </div>
                </span>
            </li>
            <li class="topbarhaslogin" style="display:none;">
                <span class="tnavsel">
                    <a href="http://i.eastmoney.com"><em class="top_name">我的股吧</em></a><em class="tnavselic"></em>
                    <div class="topnavdown">
                        <!-- <div class="topnavdownt">设置<em class="tnavselid"></em></div> -->
                        <ul class="topnavdownul">
                            <li><a href="http://i.eastmoney.com">我关注的股</a></li>
                            <li><a href="http://i.eastmoney.com">我关注的人</a></li>
                            <li><a href="http://i.eastmoney.com/myarts">我的发言</a></li>
                            <li><a href="http://i.eastmoney.com/collection">我的收藏</a></li>
                        </ul>
                    </div>
                </span>
            </li>

            <li class="topbarhaslogin" style="display:none;">
                <span class="tnavsel">
                    <em class="top_name my_msg">我的消息<i></i></em><em class="tnavselic"></em>
                    <div class="topnavdown">
                        <!-- <div class="topnavdownt">消息<em class="tnavselid"></em></div> -->
                        <ul class="topnavdownul topnavdownulmsgul">
                            <li><a href="http://i.eastmoney.com">我关注的股</a></li>
                            <li><a href="http://i.eastmoney.com">我关注的人</a></li>
                            <li><a href="http://i.eastmoney.com/myarts">我的发言</a></li>
                            <li><a href="http://i.eastmoney.com/collection">我的收藏</a></li>
                            <li><a id="my_wdmsg" href="http://i.eastmoney.com/myinfo">&nbsp;&nbsp;查看问答消息<em></em></a></li>

                        </ul>
                    </div>
                </span>
            </li>          
        </ul>
    </div>
</div>
<script>
    window.shimingOption = 2;
</script>
    <script src="//gbfek.dfcfw.com/project/guba/modules/h5Adaptation.js?r=6"></script>

    <link rel="stylesheet" href="//gbfek.dfcfw.com//project/xeditor/editor_frame.css?r=72" />

<script>
    var topictype = "20";
    var topicid = "802527008";

    var barcode = "gssz";
    var code = "gssz";
    var tabtype = "";
    var stockname = "股市实战";
    var RelatedCode = "";
    var RelatedName = "";
    var OuterCode = "gssz";
    var Category = "201";
    var Division = "2";
    var Market = "-1";
    var Type = "-1";
    var QuoteCode = "";
    var BarType = "3";
    var CodeWithMarket = "gssz";
    var qq_code = "";
    var qq_code2 = "";
    var IsStockA = "False";

    var page="1";
</script>

<style>
    #topnav li .red_point_jj {
        position: absolute;
        top: -3px;
        left: 30px;
        background-image: url(http://gbfek.dfcfw.com/gubav5/images/header_bg.png);
        background-position: -25px 0;
        width: 6px;
        height: 6px;
        margin: 0 auto;
    }

    #topnav li {
        display: inline-block;
        padding: 0 6px;
        *display: inline;
        *zoom: 1;
        position: relative;
        height: 30px;
    }

    .guba_topic .topics {
        display: none;
    }

    .guba_topic ul li {
        display: block;
        background-color: #08417f;
        width: 61px;
        height: 33px;
        margin-top: 2px;
        line-height: 33px;
        text-align: center;
    }

    .guba_topic:hover ul {
        display: block;
        position: absolute;
        top: 25px;
        left: -13px;
    }
</style>



<div id="header_ad" class="gbbody">
    <div class="dh10"></div>
    <iframe width="1000" height="60" frameborder="0" scrolling="no" marginwidth="0" marginheight="0" src="http://same.eastmoney.com/s?z=eastmoney&c=1553&op=1"></iframe>
</div>

<div id="headerban" class="gbbody">
    <div id="searchbar_wrap">
    </div>
    <div id="gbhri">
        
    </div>
</div>

<div class="zwheadbline gbbody">
    <div class="zwheadblineb"></div>
</div>
<div class="gbbody ">
    <div class="his_wrap clearfix">
        <div class="read_path">
            <a href="/" target="_self" class="icon icon_gubahome"></a>
            <a href="/" target="_self">股吧首页</a> &gt;
                <a href="/list,gssz.html" target="_self">股市实战吧（gssz）</a>
             &gt;
            <span class="path_news">正文</span>
        </div>
        <div id="zjfw" class="zjfw">
            <ul>
                <li class="zjfwdb">最近访问：</li>
                <li><span id="baHistory"></span></li>
            </ul>
        </div>
    </div>
    <div id="headerfloat"></div>
    <div id="mainbody">
        <div id="stockheader" class="gbbody">
            <div class="stockheader_wrap">
                <span id="stockif">
                    <span id="stockname" data-popstock="gssz" data-poptype="1">
                            <a href="/list,gssz.html">股市实战吧</a>
                    </span>
                </span>
                <span id="follow_wrap"></span>
               
            </div>
        </div>
        <div class="zwheadbline gbbody">
            <div class="zwheadblineb"></div>
        </div>
        <div class="zwheadpager">
            <div class="zwhpager"></div>
            <div class="zwhpagerr">
                    <a href="list,gssz.html">返回股市实战吧&gt;&gt;</a>
            </div>
        </div>
        <div id="zwmbti">
            <ul id="head_topic" class="head_topic">
                <li class="topic">
                    <ul><li><a href="http://acttg.eastmoney.com/pub/web_nr_gbzwy_ttz_01_01_01_0" >好文章，能赚钱！</a></li></ul>
                </li>
            </ul>

            <div id="zwmbtilr"></div>
        </div>
                    <div id="zwcontent">

                        <div id="zwcontt">
<div class='data' data-json='{"user_id":"1529085151583734","user_nickname":"券商中国","user_name":"k1529085151583734","user_v":31000,"user_type":0,"user_is_majia":false,"user_level":0,"user_first_en_name":"qszg","user_age":"1.5年","user_influ_level":8,"user_black_type":0}'></div>                                <div id="zwconttphoto">

                                        <a href="http://iguba.eastmoney.com/1529085151583734" data-poptype="2" data-popper="1529085151583734">
                                            <img src="http://avator.eastmoney.com/qface/1529085151583734/50" width="50" height="50" class="userphoto" />
                                        </a>

                                </div>
                                <div id="zwconttb">
                                    <div id="zwconttbn">
                                        <strong>
                                            <a href="http://iguba.eastmoney.com/1529085151583734"  data-popper="1529085151583734" data-poptype="1" target="_blank"><font>券商中国</font></a><input type="hidden" value="31000" /><em class="icon icon_jv_org" title="东方财富机构认证"></em>

                                        </strong>
                                            <div class="influence_wrap" data-user_age="1.5年" data-user_level="8" data-user_black_type="0" data-type="0"></div>
                                    </div>
                                    <div class="zwfbtime">发表于 2019-01-24 07:47:30 股吧网页版</div>

                                </div>

                            <div id="zwconttbtns">
                                <a href="//v2.eastmoney.com/" style="display: none;" id="dongmilink" class="graylink">V认证</a>&ensp;<a href="javascript:;" target="_self" id="jubaolink" class="graylink" style="color: #039;font-size: 14px;font-weight: bolder;">举报</a>
                            </div>
                        </div>
                            <div id="post_content" class="zwcontentmain xeditor">
                                <div id="zwconttbt">
平安银行可转债1400倍申购创纪录，市场又热了？还有1700亿银行可转债在路上                                                                    </div>
                                <div id="zwconbody">
                                    <div class="stockcodec .xeditor">
<p><img src="https://dfscdn.dfcfw.com/download/A120180903205892" class="__bg_gif " alt=""></p><p><img src="https://dfscdn.dfcfw.com/download/A120190124121383" class="" alt=""></p><p>图片来源：图虫创意</p><p>一举打破<span data-code="601818,1,1" data-type="1" class="zwstock">光大银行</span><span data-code="395021,2,0" data-type="1" class="zwstock">可转债</span>网下200倍的申购纪录，<span data-code="000001,2,0" data-type="1" class="zwstock">平安银行</span>日前发行的260亿元可转债，创出了A股迄今为止最火爆的申购纪录——80亿元的网下发行份额，申购总额突破10.75万亿元，有效申购倍数高达1400倍。</p><p>平安银行将创纪录“归功”于两大因素：“债的亮点”和“股的价值”。同时该行相关人士向券商中国记者透露，为保证平安可转债的价值被市场切实接受以促使转股，平安银行还将在今年安排包括业务条线开放日等在内的投资者关系活动，充分展示投资亮点，向资本市场传递管理层的决心和平安银行基本面。</p>1700亿元银行可转债在路上<p>自去年2月份证监会发布<span data-type="2" data-code="162717" class="zwstock">再融资</span>新规后，较定增、配股等其他资本工具发行周期短、发行风险小的可转债，逐渐成为当前市场环境下银行补充资本的重要选择。</p><p>2018年，共有5家银行130亿元可转债发行。进入到2019年，发行规模数倍扩容，除了已发行的260亿元平安银行可转债，还有1700亿元可转债在路上。其中，<span data-code="600919,1,1" data-type="1" class="zwstock">江苏银行</span>、<span data-code="601998,1,1" data-type="1" class="zwstock">中信银行</span>、<span data-code="601328,1,1" data-type="1" class="zwstock">交通银行</span>共计1200亿元可转债已经获批，而<span data-code="600000,1,1" data-type="1" class="zwstock">浦发银行</span>500亿元可转债已通过了股东大会通过。</p>平安银行可转债五大亮点<p>平安银行本次可转债共向原股东优先配售179.27亿元，占发行总量的68.95%；原股东优先配售后的余额，再按9:1的比例，网上发行7693万元，网下发行80亿元。根据发行结果，其网下发行部分申购总额突破10.75万亿元，中签率仅为0.074%左右，有效申购倍数高达1400倍以上，创出A股迄今为止最火爆的一单可转债申购。</p><p><strong>火爆原因可归结为五点：</strong></p><p>第一、债本身要具有配置意义：平银转债是18年以来仅有的三单AAA级可转债标的之一，预计在<span data-code="000832,1,1" data-type="1" class="zwstock"><span data-type="2" data-code="161826" class="zwstock">中证转债</span></span>指数中将获得10%以上的权重；</p><p>第二、债性要有安全垫：目前平安银行股价比上年每股净资产11.77元（即转股价）低14%，因此平银转债设置了较好的债性保护，债底条款等并不差于市场上存量的各家银行可转债；</p><p>第三、股价要有价值支撑：与目前市场存量及接下来将发行银行可转债相比，截止至1月23日收盘，平安银行的市净率最高，达到0.81。分析师称该行历史上正股波动率最强，并且近两年来股价高于转股价的天数也最长；</p><p>第四、大股东要倾力支持：平安集团和寿险已全额认购58%的转债，剩余的市场发行额度较少，对市场冲击较小；</p><p>第五、要敢于传递管理层决心：平安银行是目前唯一一家公开提出有望在年内实现转股的银行。</p>平安银行本身的基本面亮点<p>当然，除上述亮点外，平安银行基本面也获得市场认可。</p><p>有卖方分析师推荐称，平安银行目前的股价本身其实就相当于一个“可转债”或“看涨期权”，下有“保底期权”，上有“浮动收益”。</p><p><strong>分析师将 “保底期权”解释为平安存款、资产质量、公司金融和零售业务四大业务板块的提升：</strong></p><p><strong>存款方面，</strong>平安银行借助350亿金融债、260亿可转债和300亿二级债的发行，做到了长期资金改善的效果。此外，2018业绩快报显示，该行吸收存款余额2.13万亿元，较上年末增长6.4%。其中，个人存款余额4615.91亿元、较上年末大幅增长35.4%，企业存款余额1.67万亿元，较上年末增长0.5%；</p><p><strong>资产质量方面，</strong>该行虽然不良率1.75%较上年末略增0.05个百分点，但不良贷款偏离度97%，较上年末下降46个百分点。这意味着什么？分析师指出，过去不良是该行估值的最大拖累，现在终于有可能转变为释放利润的重要驱动；</p><p><strong>对公业务方面，</strong>该行在压降规模的同时践行“对公做精”，交易银行正成为新的利润增长点；</p><p><strong></strong></p><p></p><p><strong>零售私行及财富管理板块，</strong>正整合平安信托团队，力图缩短与其他零售业务见长的银行的追赶周期。</p><p><strong>“浮动收益”则被分析师归纳为三大业务亮点：</strong></p><p><strong>首先是继续强调“零售银行”和“科技银行”两大定位：</strong>全面提升零售收入、利润、贷款等占比，以及扩大科技在银行和集团业务的运用场景。这样的定位估值优势，将来有可能在股价估值中得到体现；</p><p><strong>其次是深化两大拐点：</strong>该行息差自2018年一季度以来逐季回升、目前排在股份行前列；继续出清不良，提升资产质量，将以前对估值的扣分项，变成加分享。 </p><p><strong></strong></p><p></p><p><strong>最后是延续两大新高：</strong>该行非息收入占比目前为5年内的最高，同时，清收也创出最高，收回已核销不良资产近100亿，间接增加利润，都能给与股价支撑。</p><p>整体经营回暖方面，平安银行2018业绩快报显示，预计2018年实现净利润248亿元，同比增长7%，比上年的2.61%高出4.39个百分点。</p><p><img src="https://dfscdn.dfcfw.com/download/A120180903207161" class="" alt=""></p><strong>百万用户都在看</strong><p>收盘后也赚钱！30分钟"躺赚时间"正式开始，最新交易数据来了，你赚到没？春节逆回购操作新攻略奉上</p><p>"定向降息"落地！央行首次开展2575亿TMLF，特殊时点有何深意，三大特点彰显央行用心，利率并轨将有新动作</p><p>新信号！方星海10天两次谈A股改革！正研究取消新股首日涨跌幅限制，科创板出台越早越好，进一步放开股指期货</p><p>断供说法又来！一线城市在淘宝拍卖房最近20天追平12月份全月？谁是断供"主力"？来看四大真相</p><p>中信收购广证回复来了！1.6万字涉十大重点，明确广证定位，着重提及人事问题，经纪人员劳动关系不变</p><p>巧了！A股全年蒸发14万亿，楼市全年卖了14万亿，也许不仅是巧合，2019买房or买股？这些资金正执着</p><p></p><p><img src="https://dfscdn.dfcfw.com/download/A120181015152104" class="" alt=""></p><p>券商中国是证券市场权威媒体《证券时报》旗下新媒体，券商中国对该平台所刊载的原创内容享有著作权，未经授权禁止转载，否则将追究相应法律责任。</p><p><img src="https://dfscdn.dfcfw.com/download/A120181227111531" class="__bg_gif " alt=""></p><p>（来源：券商中国的财富号 2019-01-23 05:31） <a href="http://caifuhao.eastmoney.com/news/20190124053117311720040" target="_blank" style="color:Red"><b>[点击查看原文]</b></a></p>
                                    </div>
                                </div>

                            </div>
                        <script>
                            try{
                                var imgs = document.getElementsByTagName('img')
                                if(imgs&&imgs.length>0){
                                    var len = imgs.length
                                    for(var i=0;i<len;++i){
                                        var img = imgs[i]                                        
                                        img.onerror = function(e){
                                            this.src='//gbfek.dfcfw.com/project/guba/images/error/error_img_160.png'
                                            this.className = this.className+" img_error"
                                        }
                                        
                                    }
                                }
                                var isfund=false;
                                if(isfund==false)
                                {
                                    (function (){
                                        var zwconbody = document.getElementById('zwconbody');
                                        var hideHeight =window.page > 1 ? 200 : 1500;
                                        if(zwconbody.offsetHeight > hideHeight){
                                            zwconbody.style.height = hideHeight+'px';
                                            zwconbody.style.overflow = 'hidden';
                                            zwconbody.style.position = 'relative';
                                            var newHtml = document.createElement('div');
                                            newHtml.className = 'readmore';
                                            newHtml.innerHTML = '<span class="readmore_text"><span class="icon icon_readmore"></span>展开阅读全文</span>'
                                            //var appendHtml = '<div class="readmore"><span class="readmore_text"><span class="icon icon_readmore"></span>展开阅读全文</span></div>'
                                            zwconbody.appendChild(newHtml);
                                        }
                                    })();
                                }
                            }catch(e){                                
                            }

                        </script>
                            <div class="social clearfix">
                                <div id="like_wrap" data-like_count="0"></div>
                                <div id="favorite_wrap"></div>
                                <div id="forward_wrap"></div>
                               
                                <div id="share_wrap"></div>
                                <div class="zwconbtnsi" id="zwconbtnsi_pl" onclick="try{gudong.stat(3,null);}catch (e){};">
                                    <em class="icon icon_reply"></em>
                                    <a href="javascript:;" target="_self">评论</a>
                                </div>
                            </div>

                        <div class="gubamobilegg2014" id="zwcontentbgg" style="display: none; text-align: center;">
                            <ul><li><span class="red"><a href="http://js5.eastmoney.com/tg.aspx?ID=2334 " ><strong>绝地反弹？资金流量揭晓主力动作</strong></a></span>&nbsp;<span class="black"><a href="https://acttg.eastmoney.com/pub/webtg_hskh_act_gbzwydt_01_01_01_1" ><strong>精准操盘的秘密在这里</strong></a></span>&nbsp;<span class="red"><a href="http://acttg.eastmoney.com/pub/web_nr_gblb_gbnrdbkcbwzl_01_01_01_1" ><strong>科创板申报企业名单</strong></a></span>&nbsp;<span class="black"><a href="http://fund.eastmoney.com/trade/yzzq.html?spm=gb.01.ht#zwf_,sc_jn,st_desc" ><strong>基金今年来最高涨超81%</strong></a></span>&nbsp;<span class="red"><a href="http://data.eastmoney.com/gzfx/" ><strong>如何锁定低估值个股？</strong></a></span></li></ul>
                        </div>  
                        <div class="wrap guba_warning">
                            <div style="color: #666666; padding: 20px 0 0 0; clear: both; line-height: 18px;">
                                郑重声明：用户在财富号/股吧/博客社区发表的所有信息（包括但不限于文字、视频、音频、数据及图表）仅代表个人观点，与本网站立场无关，不对您构成任何投资建议，据此操作风险自担。
                            </div>
                        </div>
                        <div style="position:relative;"><div id="storeply" name="storeply" style="position:absolute;top:-100px;"></div></div>
                                            
                        <div class="clear"></div>
                    </div>
                        <div class="zwbb1"></div>


       <!--耗时 46.875 毫秒 (2019/7/8 14:43:56 - 2019/7/8 14:43:56 )-->

        <div id="zwlist">

        </div>


        
        <div style="clear: both; height: 29px; border-top: 0; padding: 2px 0 8px 10px;">
            <div id="zwbtextlink">
                
            </div>
        </div>

            <div class="gbbox_comment" id="sendnewt">
                <div class="gbboxt"><div class="commenttext">评论该主题</div> <span class="appealspan"><a href="javascript:;" target="_self" id="appealuserlink">帖子不见了！怎么办？</a></span></div>
                <div class="gbboxtr"><span class="topbarlogin">作者：您目前是匿名发表 &nbsp; <a href="" id="bottomlogin" class="gubaloginlink"><strong>登录</strong></a> | <a href="" class="strlink gubareglink" id="bottomreg">5秒注册</a></span> <span class="topbarhaslogin">作者：<span class="headerusername"><span class="username"></span></span>，欢迎留言 <a href="javascript:;" class="logoutbtn" target="_self">退出</a></span><a class="publishtext" href="/list,gssz.html#sendnewt)"><strong>发表新主题</strong></a></div>
                <div class="gbboxb publishnew">
                    <form name="gbsform" method="post" action="" id="gbsform" class="gbsform">
                        <div class="mtj1 clearfix" id="yzmp">
                            <label for="" class="l tzla">内容：</label><span class="triangle_ico"></span><textarea class="gbsformt1 editor_title" id="gbtainput" placeholder="内容"></textarea>
                        </div>
                        <div class="mtj2">
                            <div class="editorfuns" id="editorfuns">
                                <a href="javascript:;" id="gbtainpubtn4" data-fun="face" target="_self"><em class="iconface"></em><span class="textface">表情</span></a>
                                
                            </div>
                            <div class="gbsformbtns">
                                <a href="http://acttg.eastmoney.com/pub/web_nr_gbzwy_bmy_01_01_01_0" >实名认证 ，即可获得价值100元大礼包！</a></li></ul>
                                <span id="gdregbtn"></span>&nbsp;
                                <button type="submit" class="gbsformi3">发  布</button>
                            </div>
                            <div class="clear"></div>
                        </div>

                    </form>
                </div>
            </div>
        <div class="post_warning" style="color: #666666; padding: 16px 0 0 0; clear: both; line-height: 20px;">           
            郑重声明：用户在社区发表的所有资料、言论等仅代表个人观点，与本网站立场无关，不对您构成任何投资建议。用户应基于自己的独立判断，自行决定证券投资并承担相应风险。<a href="https://passport2.eastmoney.com/pub/FAQ/Service.html" target="_blank">《服务条款》</a><a href="http://guba.eastmoney.com/CommitmentLetter.aspx" target="_blank">《跟帖评论自律管理承诺书》</a>
        </div>


                <div class="siderg" style="padding-top: 10px; clear: both;">
                    <iframe width="730" height="90" frameborder="0" scrolling="no" marginwidth="0" marginheight="0" src="http://same.eastmoney.com/s?z=eastmoney&c=1554&op=1"></iframe>
                </div>

    </div>
    <div id="sider"></div>
</div>
<div class="clear"></div>



<script>
    if (document.all && !window.XMLHttpRequest) {
        $(".qrli").hover(function () {
            $(".icon_qr", this).show();
        }, function () {
            $(".icon_qr", this).hide();
        });
    }
</script>
<div id="stockhqh" style="display: none;"></div>


<script>
    var post_article = {"post":{"post_id":802527008,"post_user":{"user_id":"1529085151583734","user_nickname":"券商中国","user_name":"k1529085151583734","user_v":31000,"user_type":0,"user_is_majia":false,"user_level":0,"user_first_en_name":"qszg","user_age":"1.5年","user_influ_level":8,"user_black_type":0,"user_bizflag":"2","user_bizsubflag":"003006","user_extend":{"uid":"1529085151583734","ncfh_authorid":222804,"ncfh_status":1,"ncfh_orgtag":"003006","ncfh_isdel":0,"ncfh_bigvip":1,"ncfh_qualificationinfo":"券商中国","ncfh_organizationtag":1,"ncfh_updatetime":"2019-07-07 20:43:31","pushtime":"2019-07-07 20:47:37"},"user_introduce":"致力于提供最及时的财经资讯，最专业的解读分析，覆盖宏观经济、金融机构、A股市场、上市公司、投资理财等财经领域。"},"post_guba":{"stockbar_type":10,"stockbar_code":"gssz","stockbar_inner_code":"gssz","stockbar_name":"股市实战","stockbar_market":"gssz","stockbar_quote":2,"stockbar_exchange":-1,"stockbar_external_code":"gssz"},"post_title":"平安银行可转债1400倍申购创纪录，市场又热了？还有1700亿银行可转债在路上","post_content":"<p><img src='https://dfscdn.dfcfw.com/download/A120180903205892' class='__bg_gif ' alt=''></p><p><img src='https://dfscdn.dfcfw.com/download/A120190124121383' class='' alt=''></p><p>图片来源：图虫创意</p><p>一举打破<span data-code='601818,1,1' data-type='1' class='zwstock'>光大银行</span><span data-code='395021,2,0' data-type='1' class='zwstock'>可转债</span>网下200倍的申购纪录，<span data-code='000001,2,0' data-type='1' class='zwstock'>平安银行</span>日前发行的260亿元可转债，创出了A股迄今为止最火爆的申购纪录——80亿元的网下发行份额，申购总额突破10.75万亿元，有效申购倍数高达1400倍。</p><p>平安银行将创纪录“归功”于两大因素：“债的亮点”和“股的价值”。同时该行相关人士向券商中国记者透露，为保证平安可转债的价值被市场切实接受以促使转股，平安银行还将在今年安排包括业务条线开放日等在内的投资者关系活动，充分展示投资亮点，向资本市场传递管理层的决心和平安银行基本面。</p>1700亿元银行可转债在路上<p>自去年2月份证监会发布<span data-type='2' data-code='162717' class='zwstock'>再融资</span>新规后，较定增、配股等其他资本工具发行周期短、发行风险小的可转债，逐渐成为当前市场环境下银行补充资本的重要选择。</p><p>2018年，共有5家银行130亿元可转债发行。进入到2019年，发行规模数倍扩容，除了已发行的260亿元平安银行可转债，还有1700亿元可转债在路上。其中，<span data-code='600919,1,1' data-type='1' class='zwstock'>江苏银行</span>、<span data-code='601998,1,1' data-type='1' class='zwstock'>中信银行</span>、<span data-code='601328,1,1' data-type='1' class='zwstock'>交通银行</span>共计1200亿元可转债已经获批，而<span data-code='600000,1,1' data-type='1' class='zwstock'>浦发银行</span>500亿元可转债已通过了股东大会通过。</p>平安银行可转债五大亮点<p>平安银行本次可转债共向原股东优先配售179.27亿元，占发行总量的68.95%；原股东优先配售后的余额，再按9:1的比例，网上发行7693万元，网下发行80亿元。根据发行结果，其网下发行部分申购总额突破10.75万亿元，中签率仅为0.074%左右，有效申购倍数高达1400倍以上，创出A股迄今为止最火爆的一单可转债申购。</p><p><strong>火爆原因可归结为五点：</strong></p><p>第一、债本身要具有配置意义：平银转债是18年以来仅有的三单AAA级可转债标的之一，预计在<span data-code='000832,1,1' data-type='1' class='zwstock'><span data-type='2' data-code='161826' class='zwstock'>中证转债</span></span>指数中将获得10%以上的权重；</p><p>第二、债性要有安全垫：目前平安银行股价比上年每股净资产11.77元（即转股价）低14%，因此平银转债设置了较好的债性保护，债底条款等并不差于市场上存量的各家银行可转债；</p><p>第三、股价要有价值支撑：与目前市场存量及接下来将发行银行可转债相比，截止至1月23日收盘，平安银行的市净率最高，达到0.81。分析师称该行历史上正股波动率最强，并且近两年来股价高于转股价的天数也最长；</p><p>第四、大股东要倾力支持：平安集团和寿险已全额认购58%的转债，剩余的市场发行额度较少，对市场冲击较小；</p><p>第五、要敢于传递管理层决心：平安银行是目前唯一一家公开提出有望在年内实现转股的银行。</p>平安银行本身的基本面亮点<p>当然，除上述亮点外，平安银行基本面也获得市场认可。</p><p>有卖方分析师推荐称，平安银行目前的股价本身其实就相当于一个“可转债”或“看涨期权”，下有“保底期权”，上有“浮动收益”。</p><p><strong>分析师将 “保底期权”解释为平安存款、资产质量、公司金融和零售业务四大业务板块的提升：</strong></p><p><strong>存款方面，</strong>平安银行借助350亿金融债、260亿可转债和300亿二级债的发行，做到了长期资金改善的效果。此外，2018业绩快报显示，该行吸收存款余额2.13万亿元，较上年末增长6.4%。其中，个人存款余额4615.91亿元、较上年末大幅增长35.4%，企业存款余额1.67万亿元，较上年末增长0.5%；</p><p><strong>资产质量方面，</strong>该行虽然不良率1.75%较上年末略增0.05个百分点，但不良贷款偏离度97%，较上年末下降46个百分点。这意味着什么？分析师指出，过去不良是该行估值的最大拖累，现在终于有可能转变为释放利润的重要驱动；</p><p><strong>对公业务方面，</strong>该行在压降规模的同时践行“对公做精”，交易银行正成为新的利润增长点；</p><p><strong></strong></p><p></p><p><strong>零售私行及财富管理板块，</strong>正整合平安信托团队，力图缩短与其他零售业务见长的银行的追赶周期。</p><p><strong>“浮动收益”则被分析师归纳为三大业务亮点：</strong></p><p><strong>首先是继续强调“零售银行”和“科技银行”两大定位：</strong>全面提升零售收入、利润、贷款等占比，以及扩大科技在银行和集团业务的运用场景。这样的定位估值优势，将来有可能在股价估值中得到体现；</p><p><strong>其次是深化两大拐点：</strong>该行息差自2018年一季度以来逐季回升、目前排在股份行前列；继续出清不良，提升资产质量，将以前对估值的扣分项，变成加分享。 </p><p><strong></strong></p><p></p><p><strong>最后是延续两大新高：</strong>该行非息收入占比目前为5年内的最高，同时，清收也创出最高，收回已核销不良资产近100亿，间接增加利润，都能给与股价支撑。</p><p>整体经营回暖方面，平安银行2018业绩快报显示，预计2018年实现净利润248亿元，同比增长7%，比上年的2.61%高出4.39个百分点。</p><p><img src='https://dfscdn.dfcfw.com/download/A120180903207161' class='' alt=''></p><strong>百万用户都在看</strong><p>收盘后也赚钱！30分钟'躺赚时间'正式开始，最新交易数据来了，你赚到没？春节逆回购操作新攻略奉上</p><p>'定向降息'落地！央行首次开展2575亿TMLF，特殊时点有何深意，三大特点彰显央行用心，利率并轨将有新动作</p><p>新信号！方星海10天两次谈A股改革！正研究取消新股首日涨跌幅限制，科创板出台越早越好，进一步放开股指期货</p><p>断供说法又来！一线城市在淘宝拍卖房最近20天追平12月份全月？谁是断供'主力'？来看四大真相</p><p>中信收购广证回复来了！1.6万字涉十大重点，明确广证定位，着重提及人事问题，经纪人员劳动关系不变</p><p>巧了！A股全年蒸发14万亿，楼市全年卖了14万亿，也许不仅是巧合，2019买房or买股？这些资金正执着</p><p></p><p><img src='https://dfscdn.dfcfw.com/download/A120181015152104' class='' alt=''></p><p>券商中国是证券市场权威媒体《证券时报》旗下新媒体，券商中国对该平台所刊载的原创内容享有著作权，未经授权禁止转载，否则将追究相应法律责任。</p><p><img src='https://dfscdn.dfcfw.com/download/A120181227111531' class='__bg_gif ' alt=''></p><p>（来源：券商中国的财富号 2019-01-23 05:31） <a href='http://caifuhao.eastmoney.com/news/20190124053117311720040' target='_blank' style='color:Red'><b>[点击查看原文]</b></a></p>","post_abstract":"图片来源：图虫创意一举打破光大银行可转债网下200倍的申购纪录，平安银行日前发行的260亿元可转债，创出了A股迄今为止最火爆的申购纪录——80亿元的网下发行份额，申购总额突破10.75万亿元，有效申购倍数高达1400倍。平安银行将创纪录“归功”于两大因素：“债的亮点”和“股的价值”。同时该行相关人士向券商中国记者透露，为保证平安可转债的价值被市场切实接受以促使转股，平安银行还将在今年安排包括业务条...","post_publish_time":"2019-01-24 07:47:30","post_last_time":"2019-01-24 07:47:30","post_display_time":"2019-01-24 07:47:30","post_ip":"","post_state":0,"post_checkState":0,"post_click_count":6608,"post_forward_count":0,"post_comment_count":0,"post_comment_authority":0,"post_like_count":0,"post_is_like":false,"post_is_collected":false,"post_type":20,"post_source_id":"20190124053117311720040","post_top_status":0,"post_status":0,"post_from":"股吧网页版","post_from_num":20,"post_pdf_url":"","post_has_pic":false,"has_pic_not_include_content":false,"post_pic_url":[],"source_post_id":0,"source_post_state":0,"source_post_user_id":"","source_post_user_nickname":"","source_post_user_type":0,"source_post_user_is_majia":false,"source_post_pic_url":[],"source_post_title":"","source_post_content":"","source_post_ip":"","source_post_type":0,"source_post_guba":{"stockbar_type":0,"stockbar_code":"","stockbar_inner_code":null,"stockbar_name":"","stockbar_market":"","stockbar_quote":-1,"stockbar_exchange":-1,"stockbar_external_code":""},"post_video_url":null,"source_post_video_url":null,"source_post_source_id":"","code_name":"gssz","product_type":"0","v_user_code":"0","source_click_count":null,"source_comment_count":"","source_forward_count":"","source_publish_time":"","source_user_is_majia":"False","ask_chairman_state":null,"selected_post_code":"","selected_post_name":"","selected_relate_guba":null,"ask_question":null,"ask_answer":null,"qa":null,"fp_code":"u29","codepost_count":1708463,"extend":null,"post_pic_url2":[],"source_post_pic_url2":[],"relate_topic":{"id":"","name":"","h5_url":"","btype":"","stype":"","guide":""},"source_extend":null,"digest_type":1,"source_post_atuser":[],"post_inshare_count":0},"rc":1,"me":"操作成功"};
    
    var is_test=false;
    var is_fake=false;
    var comment_count="0";
</script>


<script src="//gbfek.dfcfw.com/libs/jquery/1.8.3/jquery.min.js"></script>
<script src="//emcharts.dfcfw.com/suggest/stocksuggest2017.min.js"></script>


<script src="//emcharts.dfcfw.com/ec/3.10.0/emcharts.min.js"></script>


            <script src="//gbfek.dfcfw.com/gubav5/js/gubabase_da218adfac.js"></script>
            <script src="//gbfek.dfcfw.com/gubav5/js/module_0ff83f4823.js"></script>
            <script src="//gbfek.dfcfw.com/project/guba/old/news_m.js?r=72"></script>
            <script src="//gbfek.dfcfw.com/project/guba/old/news_c.js?r=72"></script>
            <script src="//gbfek.dfcfw.com/project/guba/release/news.js?r=72"></script>

<script src="//gbfek.dfcfw.com/gubav5/modules/rightpromotion/rightpromotion.js?r=5"></script>
<script src="//gbfek.dfcfw.com/tg/EMBottomSearchTG/EMBottomSearchTG-1.0.3.min.js?v=1"></script>

<script type="text/javascript">
    var newEMBottomSearchTG = new EMBottomSearchTG({
        main_width: 1000,
        main_img: 'http://z1.dfcfw.com/2015/6/10/201506101029321460568336.jpg', //推广图片  imgurl
        link: 'http://stattg.eastmoney.com/10267' //推广链接  aurl
    });
    newEMBottomSearchTG.show();
</script>




<script type="text/javascript" src="//emcharts.dfcfw.com/newsts/newsts.min.js" charset="utf-8"></script>
<script type="text/javascript" src="//emcharts.dfcfw.com/usercollect/usercollect.min.js" charset="utf-8"></script>

<script>var _gglist = 'gssz.2.股市实战.3';</script>

<script type="text/javascript" charset="UTF-8">
    try {
        var emtj_isUpload = 1;
        var emtj_pageId = 117001300541;
        var emtj_startTime = new Date().getTime();
        var emtj_logSet = "1100000000";
        var emtj_sampleRate = 100;
    } catch (err) {       
    }
</script>
<script type="text/javascript" src="https://bdstatics.eastmoney.com/web/prd/jump_tracker.js" charset="UTF-8"></script>

    
    <!-- footer-2016 -->
<style>
    .footer2016 ul {list-style: none;margin: 0;padding: 0;}
    .footer2016 a:hover{color:#ff4901 !important;text-decoration: underline !important;}
    .footer2016 .icon, .footer2016 .navlist li a {display: inline-block;}
    .footer2016 .icon {background-image: url(https://g1.dfcfw.com/g3/201802/20180206095814.png);background-repeat: no-repeat;}
    .footer2016 {margin: 0 auto;clear: both;width: 1000px;line-height: 1.1;_line-height: 1.2;font-family: simsun;font-size: 12px;border-top: 2px solid #2F5895;}
    .footer2016 .footertg {background-color: #F3F3F3;float: left;width: 100%;padding: 12px 0 0;height: 140px;}
    .footer2016 .footertg a:link, .footer2016 .footertg a:visited {color: #676767;text-decoration: none;}
    .footer2016 .qr {float: left;width: 80px;height: 140px;line-height: 150%;padding: 0 10px 0 8px;}
    .footer2016 .qr .t {font-weight: 700;font-size: 14px;padding-bottom: 10px;}
    .footer2016 .icon_qrem80, .footer2016 .icon_qrjj80 {background-image: url(https://g1.dfcfw.com/g2/201607/20160728133707.png);width: 80px;height: 80px;}
    .footer2016 .icon_qrem80 {background-position: 0 -330px;}
    .footer2016 .icon_emwxqr, .footer2016 .icon_jjwxqr {background-image: url(https://g1.dfcfw.com/g2/201607/20160728133707.png);width: 86px;height: 106px;}
    .footer2016 .icon_emwxqr {background-position: -90px 0;}
    .footer2016 .ftglist ul li.qrli {position: relative;}
    .footer2016 .ftglist ul li.qrli:hover .icon_qr, .footer2016 .scl-news .name a {display: block;}
    .footer2016 .icon_emwxqr, .footer2016 .icon_jjwxqr {display: none;position: absolute;left: 0;top: 24px;}
    .footer2016 .ftglist {float: left;height: 140px;padding: 0px 10px;}
    .footer2016 .ftglistt {font-size: 14px;font-weight: 700;line-height: 130%;padding-bottom: 6px;text-align: left;}
    .footer2016 .ftglist ul li {padding: 5px 0;color: #676767;text-align: left;}
    .footer2016 .icon_wb {background-position: -24px 0;width: 14px;height: 12px;}
    .footer2016 .icon_wx {background-position: 0 0;width: 14px;height: 11px;}
    .footer2016 .icon_note {background-position: -47px 0;width: 13px;height: 11px;}
    .footer2016 .ftglist .icon {vertical-align: -1px;margin-right: 2px;}
    .footer2016 .footertg a:link, .footer2016 .footertg a:visited {color: #676767;text-decoration: none;}
    .footer2016 .qrl {border-left: 1px solid #ddd;margin-left: 10px;padding-left: 20px;}
    .footer2016 .icon_qrjj80 {background-position: -90px -330px;}
    .footer2016 .footercr {clear: both;text-align: center;line-height: 26px;font-size: 12px;width:1000px;margin: 0 auto;height:26px;overflow:hidden;}
    .footer2016 .icon_icp {background-position: -24px -21px;width: 15px;height: 17px;}
    .footer2016 .footercr .icon {vertical-align: -3px;margin-right: 2px;}
    .footer2016 .icon_pol {background-position: 0 -22px;width: 18px;height: 20px;}
    .footer2016 .footerlinks {background-color: #2F5895;height: 30px;text-align: center;line-height: 30px;}
    .footer2016 .footerlinks a:link, .footer2016 .footerlinks a:visited, .footer2016 .footerlinks a:hover {color: #fff !important;text-decoration: none;margin: 0 14px;}
    .footer2016 .footerbz {text-align: center;padding: 12px 0;}
    .footer2016 .footerbz a {margin: 0 4px;}
    .footer2016 .footer-police {background-position: 0 -112px;width: 110px;height: 40px;}
    .footer2016 .footer-zx110 {background-position: -120px -113px;width: 110px;height: 40px;}
    .footer2016 .footer-shjubao {background-position: 0 -172px;width: 110px;height: 40px;}
    .footer2016 .footer-cxzx {background-position: -120px -174px;width: 40px;height: 40px;}
    .footer2016 .footer-shgs {background-position: -180px -174px;width: 47px;height: 40px;}
    .footer2016 .footer-12377 {background-position: 0 -54px;width: 186px;height: 40px;}
    .footer2016 .footer-yhjb {background-position: 0 -229px;width: 123px;height: 40px;}
    .footer2016 .footer-qrlast {width: 90px;}
    .footer2016 .footer-icon-qihuoqr { background: url(https://emcharts.dfcfw.com/images/qihuoqr.png);width: 80px;height: 80px;display: inline-block;background-repeat: no-repeat;vertical-align: -5px;}
</style>

<div class="footer2016">
    <div class="footertg">
        <div class="qr">
            <div class="t"><a href="http://acttg.eastmoney.com/pub/web_dfcfsy_dbtg_wzl_01_01_01_1" style="color:#000" target="_blank">东方财富</a></div>
            <a href="http://acttg.eastmoney.com/pub/web_app_dcsy_2wm_01_01_01_0" target="_blank"><em class="icon icon_qrem80"></em></a><br />扫一扫下载APP
        </div>

        <div class="ftglist">
            <div class="ftglistt">东方财富产品</div>
            <ul>
                <li><a href="https://acttg.eastmoney.com/pub/pctg_hskh_act_dfcfwmfb_01_01_01_0" target="_blank">东方财富免费版</a></li>
                <li><a href="https://acttg.eastmoney.com/pub/pctg_hskh_act_dfcfwl2_01_01_01_0" target="_blank">东方财富Level-2</a></li>
                <li><a href="http://js5.eastmoney.com/tg.aspx?ID=1483" target="_blank">投资大师</a></li>
                <li><a href="http://js5.eastmoney.com/tg.aspx?ID=2749" target="_blank">Choice金融终端</a></li>
                <li style="display: none;"><a href="http://www.langke.com/" target="_blank">浪客 - 财经视频</a></li>
            </ul>
        </div>
        <div class="ftglist">
            <div class="ftglistt">证券交易</div>
            <ul>
                <li><a href="http://gpkh.eastmoney.com/310100" target="_blank">东方财富证券开户</a></li>
                <li><a href="https://jy.xzsec.com/Trade/Buy" target="_blank">东方财富在线交易</a></li>
				<li><a href="https://acttg.eastmoney.com/pub/pctg_hskh_act_dfcfzqjy_01_01_01_0"  target="_blank">东方财富证券交易</a></li>
                <li><a href="http://www.emsec.hk/3m-Account.html" target="_blank">港美股开户</a></li>
                <li><a href="https://jy.xzsec.com/?mode=hk" target="_blank">港美股交易 </a></li>
            </ul>
        </div>
        <div class="ftglist">
            <div class="ftglistt">关注东方财富</div>
            <ul>
                <li><a href="http://weibo.com/dfcfw" target="_blank"><em class="icon icon_wb"></em>东方财富网微博</a></li>
                <li class="qrli"><a href="javascript:;" target="_self"><em class="icon icon_wx"></em>东方财富网微信</a><em class="icon icon_qr icon_emwxqr"></em></li>
                <li><a href="http://corp.eastmoney.com/Lianxi_liuyan.asp" target="_blank"><em class="icon icon_note"></em>意见与建议</a></li>
            </ul>
        </div>
        <div class="qr qrl">
            <div class="t"><a href="http://acttg.eastmoney.com/pub/web_ttjjsy_dbtg_wzl_01_01_01_1" style="color:#000" target="_blank">天天基金</a></div>
            <a href="http://js1.eastmoney.com/tg.aspx?ID=4672" target="_blank"><em class="icon icon_qrjj80"></em></a><br />扫一扫下载APP
        </div>
        <div class="ftglist">
            <div class="ftglistt">基金交易</div>
            <ul>
                <li><a href="https://trade6.1234567.com.cn/reg/step1" target="_blank">基金开户</a></li>
                <li><a href="https://trade.1234567.com.cn/login" target="_blank">基金交易</a></li>
                <li><a href="http://huoqibao.1234567.com.cn/" target="_blank">活期宝</a></li>
                <li><a href="http://fund.eastmoney.com/trade/default.html" target="_blank">基金产品</a></li>
                <li><a href="http://fund.eastmoney.com/gslc/" target="_blank">稳健理财</a></li>
            </ul>
        </div>
        <div class="ftglist">
            <div class="ftglistt">关注天天基金</div>
            <ul>
                <li><a href="http://weibo.com/ttfund" target="_blank"><em class="icon icon_wb"></em>天天基金网微博</a></li>
                <li class="qrli"><a href="javascript:;" target="_self"><em class="icon icon_wx"></em>天天基金网微信<em class="icon icon_qr icon_jjwxqr"></em></a></li>
            </ul>
        </div>
        <div class="qr qrl footer-qrlast">
            <div class="t"><a href="http://acttg.eastmoney.com/pub/web_dfcfqhsy_dbtg_wzl_01_01_01_1" style="color:#000" target="_blank">东方财富期货</a></div>
            <a href="http://acttg.eastmoney.com/pub/web_kh_dcsy_dibudfcfqh_01_01_01_1" target="_blank"><em class="footer-icon-qihuoqr"></em></a><br>扫一扫下载APP
        </div>
        <div class="ftglist">
            <div class="ftglistt">期货交易</div>
            <ul>
                <li><a href="http://acttg.eastmoney.com/pub/web_kh_dcsy_dibuqhsjkh_01_01_01_1" target="_blank">期货手机开户</a></li>
                <li><a href="https://acttg.eastmoney.com/pub/web_kh_dcsy_dibudnkh_01_01_01_1" target="_blank">期货电脑开户</a></li>
                <li><a href="http://acttg.eastmoney.com/pub/web_kh_dcsy_dibuqhgfwz_01_01_01_1" target="_blank">期货官方网站</a></li>
            </ul>
        </div>
    </div>
        <div class="footercr" style="padding-top:8px;">信息网络传播视听节目许可证：0908328号 经营证券期货业务许可证编号：913101046312860336 违法和不良信息举报:021-34289898 举报邮箱：<a target="_self" href="mailto:jubao@eastmoney.com">jubao@eastmoney.com</a></div>
        <div class="footercr" style="padding-bottom:8px;">
            <em class="icon icon_icp"></em>沪ICP证:沪B2-20070217 <a target="_blank" rel="nofollow" href="http://shcainfo.miitbeian.gov.cn" style="color: #3F3F3F;text-decoration:none;">网站备案号:沪ICP备05006054号-11 </a> <a target="_blank" rel="nofollow" href="http://www.beian.gov.cn/portal/registerSystemInfo?recordcode=31010402000120" target="_blank" style="color: #3F3F3F;text-decoration:none;"><em class="icon icon_pol"></em>沪公网安备 31010402000120号</a> 版权所有:东方财富网 意见与建议:021-54509966/021-24099099
        </div>
    <div class="footerlinks">
			<a href="http://about.eastmoney.com" target="_blank" rel="nofollow">关于我们</a>
			<a href="http://emhd2.eastmoney.com/market" target="_blank" rel="nofollow">广告服务</a>
			<a href="http://about.eastmoney.com/home/contact" target="_blank" rel="nofollow">联系我们</a>
			<a href="http://eastmoney.zhiye.com" target="_blank" rel="nofollow">诚聘英才</a>
			<a href="http://about.eastmoney.com/home/disclaimer" target="_blank" rel="nofollow">免责声明</a>
			<a href="http://about.eastmoney.com/home/legal" target="_blank" rel="nofollow">法律声明</a>
            <a href="http://about.eastmoney.com/home/conceal" target="_blank" rel="nofollow">隐私保护</a>
			<a href="http://about.eastmoney.com/home/parper" target="_blank" rel="nofollow">征稿启事</a>
			<a href="http://www.eastmoney.com/sitemap.html" target="_blank">网站地图</a>
			<a href="http://www.zx110.org/cxs/index.html" target="_blank" rel="nofollow">放心搜</a>
			<a href="http://about.eastmoney.com/home/links" target="_blank" rel="nofollow">友情链接</a>
    </div>
    <div class="footerbz">
        <img src="https://g1.dfcfw.com/g3/201905/20190531140719.png" title="亲爱的市民朋友，上海警方反诈劝阻电
话“962110”系专门针对避免您财产被
骗受损而设，请您一旦收到来电，立即
接听。" style="vertical-align: bottom;">
        <a rel="nofollow" href="http://www.cyberpolice.cn/" class="icon footer-police" title="上海网警网络110" target="_blank"></a>
        <a rel="nofollow" href="http://www.zx110.org/" class="icon footer-zx110" title="网络社会征信网" target="_blank"></a>
        <a rel="nofollow" href="http://www.shjbzx.cn/" class="icon footer-shjubao" title="上海违法和违规信息举报中心" target="_blank"></a>
        <a rel="nofollow" href="http://www.51315.cn/company_details_1138" class="icon footer-cxzx" title="诚信在线" target="_blank"></a>
        <a rel="nofollow" href="https://www.sgs.gov.cn/lz/licenseLink.do?method=licenceView&amp;entyId=dov73ne2cwgd88longlpxsn0hcrfn5m2aa" class="icon footer-shgs" title="上海工商" target="_blank"></a>
        <a rel="nofollow" href="http://www.12377.cn" class="icon footer-12377" title="中国互联网违法和不良信息举报中心" target="_blank"></a>
        <a rel="nofollow" href="http://report.12377.cn:13225/toreportinputNormal_anis.do" class="icon footer-yhjb" title="上海市互联网违法和不良信息举报中心" target="_blank"></a>
	</div>
</div>

<script>
    if(document.all && !window.XMLHttpRequest){
        $(".qrli").hover(function(){
            $(".icon_qr",this).show();
        },function(){
            $(".icon_qr",this).hide();
        });
    }
</script>
</body>
</html>

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