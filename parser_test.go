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

<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>多用户商城系统：钢铁电商生存者与出局者的竞争 _HiMall</title>
<meta name="keywords" content="多用户商城,多用户商城系统" />
<meta name="description" content="HiShop最新消息， 经常有朋友与笔者聊起关于钢铁电商未来竞争格局的话题，听到最多并带有总结性的结论是：整个产业链上将仅有一两个既强又大的平台生存得很好，而排在三、四名以" />
<meta http-equiv="Cache-Control" content="no-transform"/>
<meta http-equiv="Cache-Control" content="no-siteapp" />
<meta http-equiv="mobile-agent" content="format=xhtml;url=https://m.hishop.com.cn/himall/show_46945.html">
<link rel="alternate" media="only screen and(max-width: 640px)" href="https://m.hishop.com.cn/himall/show_46945.html">
<meta name="applicable-device" content="pc">
<link rel="stylesheet" type="text/css" href="/templets/default/style/head_footer_2017.css"/>
<link rel="stylesheet" href="/templets/default/style/list_article_2015.css" type="text/css"/>
<script type="text/javascript" src="/templets/default/js/index/jquery-1.7.2.min.js"></script>
<script src="/scripts/wap.js" type="text/javascript"></script>
<script src="/templets/default/js/art_wap.js" type="text/javascript" charset="utf-8"></script>
</head>

<body>
<div class="head clearfix">
	<div class="head_top">
		<div class="width">
			<p>商城类电商软件及服务提供商--HiShop海商，欢迎您！</p>
			<a href="http://www.hishop.cn/" target="_blank" rel="nofollow">北京公司</a>|
			<a href="http://www.hishop.net/" target="_blank" rel="nofollow">深圳公司</a>|
			<span><script src='/plus/rpad/js.php?aid=117' language='javascript'></script></span>
		</div>
	</div>
	<div class="head_menu width clearfix">
		<a href="/" class="logo"><img alt="多用户商城,多用户商城系统" src="/templets/default/images/2017/top_logo.jpg"/></a>
		<p>
			<a id="head_zc" href="http://member.hishop.com.cn/Register.aspx" target="_blank" rel="nofollow" class="reg">注册有礼</a>
			<a href="http://member.hishop.com.cn/login.aspx?/himall/show_46945.html" target="_blank" rel="nofollow" class="login">登录</a>			
		</p>
		<ul class="nav">
			<li><a href="/">首页</a></li>
			<li><a href="/products/himall/">HiMall</a>			
			</li>
			<li><a href="/products/himall/banben/">同城O2O系统</a>				
			</li>
			<li><a href="/products/himall/himall_product.html">产品功能</a></li>
			<li><a href="/products/himall/app/">商家管理APP</a></li>
			<li><a href="/products/himall/himall_casepage.html">成功案例</a></li>
			<li><a href="/products/himall/b2b/">B2B功能</a></li>
		</ul>
	</div>
</div>	
  <div class="banner clear" style="background:url(/templets/default/hiscxt/images/arc_himall.jpg) no-repeat top center;">
<a class="banner_a"  onclick="popup('banner_试用','申请多用户商城系统试用帐号，抢先体验')"> </a></div> 

<div class="w1200">
<div class="main_left f_left">
	<div class="floor1 ">          
    	<div class="title f_left"><a href='http://www.hishop.com.cn/'>HiShop</a> > <a href='/himall/'>多用户商城系统</a> > <a href='/himall/b2b2c/'>b2b2c运营</a> > </div>
    </div>
   	
    <div class="content">
    	<div class="title">
            <h1>多用户商城系统：钢铁电商生存者与出局者的竞争 </h1>
            <div class="release">           
            <em>2017-12-06</em>|<em>HiShop</em>|<em>阅读量：<script src="/plus/count.php?view=yes&aid=46945&mid=21669" type='text/javascript' language="javascript"></script></em>
            </div>
        </div>
        
		<div class="takeaway"><b>导读：</b>HiShop最新消息， 经常有朋友与笔者聊起关于钢铁电商未来竞争格局的话题，听到最多并带有总结性的结论是：整个产业链上将仅有一两个既强又大的平台生存得很好，而排在三、四名以...</div>
		
        
    	<div class="article">HiShop最新消息，经常有朋友与笔者聊起关于钢铁电商未来竞争格局的话题，听到最多并带有总结性的结论是：整个产业链上将仅有一两个既强又大的平台生存得很好，而排在三、四名以后的众多平台将在市场竞争中成为出局者。
<p>
	笔者认为，如今的钢铁电商领域，虽然已经进入了数百个平台竞相角斗的&ldquo;春秋战国时期&rdquo;，但未来不一定只有前面所说的那一条路可选择&mdash;&mdash;只有一两个平台垄断市场。尽管人们有&ldquo;合久必分、分久必合&rdquo;之说，但这样的分和合不应该只停留在平台的存量上，而应该更多地体现在质的提升上。因为谈到平台的定义，泛泛地说，每一台手机都可以是一个接收发送各类信息、完成交易的平台，故而仅仅从量上预测未来的竞争格局未免有失偏颇。为此，笔者从现实出发，对未来钢铁电商竞争格局作一些试探性的分析和预测。</p>
<p style="text-align: center;">
	<img alt='多用户商城系统：钢铁电商生存者与出局者的竞争 ' src="/uploads/allimg/171206/21669-1G206135US16.jpg" style="width: 397px; height: 294px;" /></p>
<p>
	<strong>大小平台齐发力，相互融合谋发展</strong></p>
<p>
	商道兴衰从来就与&ldquo;天时地利人和&rdquo;密切相关。</p>
<p>
	天时方面，党的十九大提出构建人类命运共同体，实现共赢共享。在这一大背景下，作为中国钢铁产业链上的钢铁人理应立足本职，胸怀大局，撸起袖子，脚踏实地干一番事业。从哪里着手？就从构建一条完善的承载着钢铁大国梦的我国钢铁产业链做起。目前，我国钢材市场还存在着诸多不尽人意的地方。针对一些问题，国家相关部门已经出台进一步深化改革的各项激励政策，在新长征路上，改革大有可为。</p>
<p>
	地利方面，互联网技术为我们整合资源，实现互利共赢和弯道超车，打开了方便之门。下一步要做的是，利用云计算、大数据技术，建立起市场诚信经营体系，整合钢铁产业链上中下游各类有效资源，促使大小平台融合，线上虚拟市场和实体现货市场融合，平台与包括物流、加工企业在内的上中下游企业融合。整合钢铁产业链内外一切可以整合到的有效资源，营造出一个与传统市场不一样的经营氛围，建立钢铁生产、经营新秩序。</p>
<p>
	人和方面，钢铁强国梦激发了中国钢铁人无比高涨的热情，树立以和为贵的理念,这也可视为是一种文化自信的表现。近年来，众多平台远离恶性竞争，联手组织各种论坛，进行多样性的合作。如今，不仅有钢贸商参与平台交易，而且有越来越多的钢厂将资源放到网上销售，很多专业仓库、物流配送加工企业也成为平台服务体系中不可或缺的一部分。这从一个侧面表明，钢铁电商市场不大可能再出现与&ldquo;天时地利人和&rdquo;背道而驰的传统竞争模式&mdash;&mdash;大鱼吃小鱼，小鱼吃虾米，最终形成一两个超级企业垄断市场的格局。相反，市场将出现这样的情况：在互利共赢的旗帜下，大小平台融合发展，不断提高服务质量，安全、高效、低成本地助力我国钢铁产业链实现现代化。</p>
<p>
	目前，有迹象表明，不少钢铁电商在布点的同时，通过结盟、互相持股、互补互助等方式做强做实平台的服务功能，提高信息、资源、金融、物流等方面的优势，扩大了交易覆盖面，增加了交易流量，为日后的融合打下了坚实的基础。今年上半年，不少钢铁电商的交易量已经达到3000多万吨，并呈现出强劲的上升势头。如此规模的平台，仅上海就有3家-4家，其流量已经达到本世纪初上海钢贸行业鼎盛时期的交易总量。最近，又有消息在业内广为流传&mdash;&mdash;不少钢铁企业与平台合作，将资源在平台上挂牌，以加大线上销售力度。</p>
<p>
	可以预见的是，至今仍在烧钱却无法盈利的平台中，有一部分终将在市场竞争中成为出局者，而一些富有特色的平台，无论目前的体量是大或小，不但不会倒下，而且会长期存在，并会在与各平台和各相关企业的互动中，利用自身的特色，不断优化服务，朝着做强做大的方向迈进。今人扼腕的是，目前这样有特色的平台还不多，希望以后会有更多。</p>
<p>
	<strong>业绩彰显于线上，功夫落实于线下</strong></p>
<p>
	在互联网技术普及之初，各平台因为缺乏对新技术应用的全面认识，偏重于编制网上流程，消耗了大量人力、物力和财力。具体表现在：有的平台在唯技术理念的指导下，多层次组建技术团队，结果还是出现线上线下脱节的状况，这又在一定程度上影响了网上流程的畅通。无奈之下，只得依靠打补丁维持运行。</p>
<p>
	据HiShop<a href="http://www.hishop.com.cn/products/himall/" target="_blank">多用户商城系统</a>了解到，有一家平台不但独自组建了一家元件开发有限公司，还组建了流程编制团队和应对日常运行故障的维修保养队伍。尽管线上流程布置得严丝合缝，但在实施过程中还是出现了线上与线下脱节的情况，致使线上交易流程不畅通。而另一家平台在组建技术团队后，提货单上仍然需要再盖3个章才能让钢材出库，远未达到安全、高效、快捷、降低成本的目的。</p>
<p>
	随着钢铁行业两化融合工作的推进，上游钢厂的生产自动化程度日渐提高，下游用钢终端用户的自动化生产水平也在不断提高，而处于中游的仓储、物流、加工环节几乎是原地踏步，维持旧时模样。因此，平台在对上、中、下游企业实施无缝连接时，应重视钢铁存储、加工、流通等中游企业在智能化方面的建设，为专业的钢铁仓储、加工、运输企业的智能化运作提供帮助。</p>
<p>
	有的平台实力雄厚、条件优越，有能力自建智能化仓库，并在同行业中提供智能化运作的样板。这种由平台自建的方式投入非常大，不适合普及推广。对于大多数平台来说，还是应该立足现实，加大与相关企业的磨合力度，运用互联网技术不断在钢铁仓储、加工、物流等环节进行创新。这就需要平台在关注线上运作的同时，还要关注线下运作情况，适时向相关参与合作的企业提出智能化运作要求。在条件具备的情况下，平台可尝试与专业钢铁物流企业共建智能化物流体系，为行业提供实践经验并加以推广。</p>
<p>
	<strong>不忘初心争回报，&ldquo;五流&rdquo;汇聚谱新篇</strong></p>
<p>
	在设置交易流程时，平台一定要不忘初心，以普惠共享的互联网理念为指导，在每个细节上让参与者体验到新技术带来的轻松欢愉。可是，在现实生活中，平台真正做到普惠共享并非易事。</p>
<p>
	前几年，笔者从一位钢铁电商平台负责人处得知，该平台在融得资金加点贷出的基础上，还将做足融贷周期的文章。经粗略计算，即使不拉大价差，平台也能在吨钢上获取一两百元的利润。这家平台如此操作了两三年后，也就关门了。如此过分的&ldquo;玩法&rdquo;，虽然在行业里并不多见，但也在一定程度上反映出某些平台忘却了初心。即便到了今天，还是有这样的情况存在。</p>
<p>
	例如，有些平台在用自有资金放贷时，非得在银行贷款利率上加点，把平台提供服务的窗口搞成盈利的节点。实际上，银行利率点中已经包含着平台自身应得的那部分利润，再加点不仅变成了&ldquo;高利贷&rdquo;，而且提高了交易过程中资金运作的成本，对贷款企业有所不公。如果平台秉承了普惠共享的理念，那么在做融贷款业务时，一定会考虑到贷款方的利益，通过创新互联网金融方法来降低线上资金运作成本，而不是在盈利基础上再加点贷出。</p>
<p>
	上面所述的只是平台&ldquo;五流&rdquo;（信息流、资源流、商流、资金流和人才流）之一的资金流的问题。笔者相信，随着深化改革的推进，平台将树立起普惠共享的服务观念，构建有效的线上线下互动系统，以达到安全、快捷、高效和降低成本的目的，伴之而来的一定是开放的、新颖的平台。这样的平台将汇聚&ldquo;五流&rdquo;，构建钢铁流通新格局。</p>
<p>
	在过去几年里，钢铁电商平台把注意力集中在增加流量上，这无可厚非。但是，随着互联网技术的普及应用，用户对新技术应用的期望值也在提高，平台在实际运作中只抓流量已经远远不够了。平台建设是一项系统工程，需要线上线下各个环节的相互协同。目前看来，多数平台在抓&ldquo;五流&rdquo;的过程中，忽视了对相关人才的培养和发现。因此，在对未来发展进行规划时，各大平台应做好人才方面的功课，一方面通过优秀人才，发挥好信息流的优势，为上、下游企业介绍国内外先进的管理理念、技术设施，另一方面也可以为上下游企业推荐所需优秀人才，使人才流成为上下游企业沟通联动的纽带。</p>
<p>
	在中游企业中，平台除了可以帮助企业共建智能化物流体系外，还应在流通商中培育为终端用户服务的类销客队伍，千方百计地把平台拥有的优势服务与流通商共享，加大为终端用户服务的力度。平台千万不要把&ldquo;精简中介环节&rdquo;中的&ldquo;去中介&rdquo;看作是去掉中间钢贸商，事实上，截至目前，线上流量中的绝大部分都是由钢贸商创造的。平台在为终端用户服务的过程中，应学会融合，帮助钢贸商不断提高其为终端用户服务的质量。</p>

                      <p>文章来源：https://www.hishop.com.cn/himall/show_46945.html</p>
<p><海商（www.hishop.com.cn）是国内知名商城系统及商城网站建设提供商，HiMall为企业级商家提供<a href="/products/himall/">B2B2C商城系统</a>、多店铺商城系统、<a href="/products/himall/">多用户商城系统</a>等多端商城及电子商务行业解决方案></p>
<p>申明：本网站部分文章和图片来源网络编辑，如有侵权及时沟通删除。海商hishop网站原创文章，转载请注明来源。</p>
               <div class="mk_03 clear">
        <ul>
           
    </ul>
           
       </div>      
         <!-- 新媒体广告 -->
            <p class="newmide" style="text-align:center;margin:10px 0;"><script src='/plus/rpad/js.php?aid=81' language='javascript'></script></p>        
       </div>
        <div class="tags">
            <div>TAGS: <a href='/tags/2675.html'>B2B2C商城系统</a><a href='/tags/2715.html'>多用户商城系统</a></div>
            <div class="Turn">
            <em>上一篇：<a href='/himall/show_46944.html'>做区域共同配送B2B经销商怎样实现盈利？ </a> </em>
            <em>下一篇：<a href='/himall/show_47027.html'>超市订货管理系统具备哪些功能优势？</a> </em>
            
            </div>
        </div>
        
        <div class="details">
        	<div class="title"><em></em>推荐阅读</div>
             <ul class="bd-l">
             
          </ul>       
        </div>        
	</div>
</div>

<div class="main_right f_right ">
  <div class="floor6 clear">
        <div class="title"><span>今日热点</span></div>
          <p class="floorp">
           最新消息， HiShop消息，做多用户商城系统的运营是需要时间和精力去做的，想要做好运营不是那么容易的事，但只要思路明确，方法科学也是完全可以做好的。那么怎么做好运营呢?那么，B2B2C 多用户商城系统 运营需要做到哪七点? 个体或者小企业在天猫、....</a>

          </p>
  </div> 
    <div class="floor6 clear">
        <div class="title"><span>热门文章</span></div>
        <ul>
            <li>
                   
                        <a href="/himall/show_71386.html" target="_blank">b2b2c多用户商城系统的知识点，谁都得看看</a>
                             <p class="time">2小时前</p>
             </li>
<li>
                   
                        <a href="/himall/show_71800.html" target="_blank">b2b2c商城系统功能之经典案例爱乐商城</a>
                             <p class="time">3天前</p>
             </li>
<li>
                   
                        <a href="/himall/show_71744.html" target="_blank">b2b2c网上商城系统怎么做可以直接提升商城的用户流量！</a>
                             <p class="time">4天前</p>
             </li>
<li>
                   
                        <a href="/himall/show_71693.html" target="_blank">怎么利用b2b2c商城系统源码来直接提升商城自身？</a>
                             <p class="time">5天前</p>
             </li>
<li>
                   
                        <a href="/himall/show_71629.html" target="_blank">对于b2b2c商城系统要多少钱不少很清楚？我们拿易订商城的例子</a>
                             <p class="time">6天前</p>
             </li>
<li>
                   
                        <a href="/himall/show_71590.html" target="_blank">多商户B2B2C商城系统怎么从用户去入手？抓住市场！</a>
                             <p class="time">7天前</p>
             </li>
<li>
                   
                        <a href="/himall/show_71545.html" target="_blank">b2b2c商城系统功能要怎么去布置？哪些要哪些不必要？</a>
                             <p class="time">10天前</p>
             </li>
<li>
                   
                        <a href="/himall/show_71501.html" target="_blank">现在的b2b2c网上商城系统有没改进的办法？</a>
                             <p class="time">11天前</p>
             </li>
  
        </ul>
    </div> 

    <div class="floor5">
    	<div class="title"><span>产品推荐</span></div>
        <ul>
        	<li >
                <div class="pic_1"><a href="/products/himall/" target="_blank">
                <img src="/templets/default/ecschool/images/temp/16_1.jpg"></a></div>                
                <em>Himall</em><br />
                <span>构筑京东，天猫同等级多用户商城系统，首个打通全网pc+微信+app+触屏的B2B2C多用户商城系统。为用户带来互联网最便捷的盈利模式，精细化运营，精准化推广，充分挖掘互联网市场潜力，点击详情，马上申请免费体验。</span><a href="/products/himall/" target="_blank">详情</a>                       
            </li>
        </ul>
    </div>
    <div class="floor7">
        <div class="title"><span>经典阅读</span></div>
    	<ul>
            <li><a href="/products/himall/">b2b2c商城系统</a><a href="/products/himall/">多商家商城系统</a></li>
            <li><a href="/himall/">多用户商城系统</a><a href="/products/himall/">多商户商城系统</a></li>
             <li><a href="/ydsc/dsxt/">电子商务系统</a><a href="/ydsc/dsxt/">电商网站建设</a></li>
            <li><a href="/himall/scym/" target="_blank">b2b2c商城源码</a><a href="/himall/scym/">b2b2c源码</a></li>
        </ul>
    </div>
  <div class="floor8 clear">
      <div class="title"><span>热门搜索</span></div>
       <div class="fl8">
       <p><a href='/products/himall/'>多用户商城系统</a></p><p><a href='/products/himall/'>B2B2C商城系统</a></p>
       <p><a href="/himall/scym/">b2b2c源码</a></p><p><a href="/himall/scym/">b2b2c商城源码</a></p>
        <p><a href='/tags/16679.html'>B2B2C商城系统功能</a></p><p><a href='/tags/18900.html'>多商户B2B2C商城系统</a></p><p><a href='/tags/13085.html'>京东新通路</a></p><p><a href='/tags/19643.html'>2019中国电商网站排名</a></p><p><a href='/tags/19628.html'>b2b2c与s2b2c</a></p><p><a href='/tags/19594.html'>b2b2c系统展示</a></p><p><a href='/tags/16685.html'>B2B2C商城系统价格</a></p><p><a href='/tags/19536.html'>社区团购平台怎么做</a></p><p><a href='/tags/19367.html'>新零售B2B2C商城系统</a></p><p><a href='/tags/19386.html'>b2b2c系统平台</a></p><p><a href='/tags/11355.html'>网易考拉</a></p><p><a href='/tags/14531.html'>商城系统开发</a></p><p><a href='/tags/18170.html'>B2B2C商城系统架构</a></p><p><a href='/tags/19366.html'>B2B2C网站管理系统</a></p><p><a href='/tags/19066.html'>B2B2C商城系统代码</a></p><p><a href='/tags/11461.html'>网易严选</a></p><p><a href='/tags/19255.html'>b2b2c电商开发</a></p><p><a href='/tags/15311.html'>B2B2C网站建设</a></p><p><a href='/tags/19226.html'>b2b2c商城功能</a></p><p><a href='/tags/19265.html'>b2b2c系统多少钱</a></p>
        </div>
    </div>
    <div class="floor9">
		<div class="title"><span>推荐专题</span></div>
		<i>
			<a href="/himall/show_65773.html" target="_blank">
					<img src="/uploads/190129/21662-1Z129110200a1.jpg"/>
					<p>电商系统源码_商城系统源码_电商网站系统源码怎么选？</p>
				</a>
<a href="/himall/show_65592.html" target="_blank">
					<img src="/uploads/allimg/190124/21716-1Z124152T60-L.jpg"/>
					<p>B2B2C网站系统-B2B2C网站系统哪个好-B2B2C商城系统</p>
				</a>
<a href="/himall/show_65746.html" target="_blank">
					<img src="/uploads/181220/21704-1Q220142523913.png"/>
					<p>多用户商城源码_多用户商城系统源码_多用户商城开发</p>
				</a>
<a href="/himall/show_65751.html" target="_blank">
					<img src="/uploads/181220/21704-1Q220142523913.png"/>
					<p>b2b2c商城多少钱_b2b2c商城系统报价_B2B2C商城开发</p>
				</a>
<a href="/himall/show_65769.html" target="_blank">
					<img src="/images/defaultpic.gif"/>
					<p>B2B2C建站_多用户建站_B2B2C多用户建站怎么搭建？</p>
				</a>
<a href="/himall/show_65784.html" target="_blank">
					<img src="/images/defaultpic.gif"/>
					<p>深圳B2B2C电商系统哪家好？就选HiMallB2B2C商城系统</p>
				</a>

		</i>
	</div>
</div>
</div>

<div class="new_footer">
	<div class="width t">
		<div class="fl">
		分公司：<a href="http://www.hishop.cn/" target="_blank" rel="nofollow">北京公司</a>|<a href="http://www.hishop.net/" target="_blank" rel="nofollow">深圳公司</a>|<a>广州公司</a>
		</div>
		<dl class="fl">
			<dt><span></span>总部电话：<em>4006-089-731</em></dt>
			<dd>
				<span></span>
				北京公司<br />
				<em>4006-601-982</em>
			</dd>
			<dd>
				<span></span>
				深圳公司<br />
				<em>4008-400-198</em>
			</dd>
			<dd>
				<span></span>
				广州公司<br />
				<em>4006-808-668</em>
			</dd>
		</dl>
		<ul class="fr">
			<li>
				<a><em></em></a>
				干货分享
			</li>
			<li>
				<img src="/templets/default/images/2017/footer_wx.png"/>
				官方微信
			</li>
		</ul>
	</div>
	<div class="width">
		<div class="s_nav">
			<a href="/hishop/about/"rel="nofollow">关于我们</a>|<a href="/hishop/buy/"rel="nofollow">如何购买</a>|<a href="/hishop/contactus/"rel="nofollow">联系我们</a>|<a href="https://help.hishop.com.cn" target="_blank" rel="nofollow">帮助中心</a>|<a href="/hishop/job/" rel="nofollow">诚聘英才</a>|<a href="/hishop/sitemap/">网站地图</a>|<a href="/hishop/link/" rel="nofollow">友情链接</a>|<a href="/hishop/legalnotices/" rel="nofollow">法律声明</a>|<a href="http://mb.hishop.com.cn/" target="_blank" rel="nofollow">商城模板</a>|<a href="/service/" rel="nofollow">技术服务</a>
		</div>
		<div class="copy">
			软件企业编号：湘R-2009-0055&nbsp;&nbsp;|&nbsp;&nbsp;高新技术企业编号：GR201043000270&nbsp;&nbsp;|&nbsp;&nbsp;增值业务电信许可证：湘B2-20160062<br />
			软件产品编号：湘DGY-2009-0155&nbsp;&nbsp;|&nbsp;&nbsp;软件著作权登记号：2006SR09906 2009SR047196<br />
			湖南省职业经理人协会理事&nbsp;&nbsp;|&nbsp;&nbsp;长沙市电子商务协会理事&nbsp;&nbsp;|&nbsp;&nbsp;长沙市软件协会会员&nbsp;&nbsp;|&nbsp;&nbsp;长沙服务外包行业协会会员<br />
			Copyright©2003-2019 HiShop商城系统<a href="/">微信分销系统</a>以及<a href="/products/vfx/">微分销系统</a>多用户B2B2C商城系统一站式网上<a href="/">商城网站建设</a>服务商All Rights Reserved.湘ICP备08105577号
			<p>
				<a href="http://www.beian.gov.cn/portal/registerSystemInfo?recordcode=43010202000038" rel="nofollow" target="_blank" class="n1"></a>
				<a href="http://www.hngswj.gov.cn/ei/Ei!certificate.action?id=8d5edfc6ad054507b73396b8b1802a5e" rel="nofollow" target="_blank" class="n2"></a>
			</p>
		</div>
	</div>
	<!--hishoppc-->
	<script>
var _hmt = _hmt || [];
(function() {
  var hm = document.createElement("script");
  hm.src = "https://hm.baidu.com/hm.js?779041bffe6533ef44f74c2fdc69a426";
  var s = document.getElementsByTagName("script")[0]; 
  s.parentNode.insertBefore(hm, s);
})();
</script>
</div>
<!--在线客服-->
<!--<script src='/plus/rpad/js.php?aid=23' language='javascript'></script>
<script  type="application/ld+json"> {
       "@context": "https://zhanzhang.baidu.com/contexts/cambrian.jsonld",
       "@id": "https://www.hishop.com.cn/himall/show_46945.html",
        "appid": "1600581270734074",
       "title":"多用户商城系统：钢铁电商生存者与出局者的竞争 ",
       "images": [
           "/uploads/allimg/171206/21669-1G206135US16-lp.jpg"        
       ],
       "description": "HiShop最新消息， 经常有朋友与笔者聊起关于钢铁电商未来竞争格局的话题，听到最多并带有总结性的结论是：整个产业链上将仅有一两个既强又大的平台生存得很好，而排在三、四名以",
       "pubDate": "2017-12-06T13:56:45",
       "puDate": "2019-07-01T11:05:45",
       "data":{
        "data.WebPage":{
          "pcUrl":"https://www.hishop.com.cn/himall/show_46945.html",
          "wapUrl":"https://m.hishop.com.cn/himall/show_46945.html"
        }
        }
}</script>-->
<!-- 自动推送 -->

<script>
(function(){
    var bp = document.createElement('script');
    var curProtocol = window.location.protocol.split(':')[0];
    if (curProtocol === 'https') {
        bp.src = 'https://zz.bdstatic.com/linksubmit/push.js';        
    }
    else {
        bp.src = 'http://push.zhanzhang.baidu.com/push.js';
    }
    var s = document.getElementsByTagName("script")[0];
    s.parentNode.insertBefore(bp, s);
})();
</script>
<script type="text/javascript">
$(document).ready(function(){
  $('a:[href="http://member.hishop.com.cn/Register.aspx"]').click(function(){
    $(this).attr("href","http://member.hishop.com.cn/Register.aspx?myurl="+window.location.href);
  })
});
</script>
<!--弹出层开始-->
	<link rel="stylesheet" type="text/css" href="/templets/default/style/popup.css"/>
    <script src="/templets/default/js/crm_ajax.js" type="text/javascript"></script>
	<div id="back_over"></div>
    <div class="popup_bg" id="pop_box" style="display:none;">
    	<img src="/templets/default/images/popbox/pop_icot.png"/>
    	<div class="pop_box">     		
	        <div class="pop_tit"><em>获取解决方案</em><span id="popup_x" onclick="popclose();" title="关闭">
				<img class="popup_close" src="/templets/default/images/popbox/pop_close.png"/></span>
	        </div>
	        	<input type="hidden" id="txtInfoCategoryId" name="InfoCategoryId" value="3" />
		        <input type="hidden" name="txtProductName" id="txtProductName" value="Himall" />  
	        	<input type="hidden" name="mudi" id="mudi" value="" />
		        <input type="hidden" name="myurl" id="myurl" value="" />
		        <input type="hidden" name="jump" id="jump" value="https://www.hishop.com.cn/register/himall_ty.html" />
		        <ul class="pop_cont">
		        	<li><input onkeyup="value=value.replace(/\D+/g,'')" id="Phone" name="Phone" type="text"  placeholder="您的手机号码" /></li>
		            <li><input name="validate" type="text" id="validate" placeholder="图形验证码" style="width:70%"/>
		            	<img src='/include/vdimgck.php?iw=112&ih=38&fz=18' onClick="this.src=this.src+'?'" /></li>
		            <li><input name="PhoneCode" id="PhoneCode" type="text" placeholder="请填写手机验证码" style="width:60%;" />
		            	<input class="input_c" type="button" value="获取手机验证码" onclick="sendPhoneSMS()" /></li>  
		            <li><input name="button" id="popup_btn" type="button" value="提  交" onclick="FormSub()" /></li>	            
		        </ul>
			<div class="tel">咨询电话：4006-089-731</div>
	    </div>
    </div>
<script type="text/javascript">
$("#myurl").val(window.location.href);
	// 弹窗
function popup(mudi,title){
        var maskWidth = $(window).width(); //窗口的宽度		
        var maskHeight = $(window).height(); //窗口高度		
        var popTop = (maskHeight / 2) - ($('#pop_box').height() / 2);
        var popLeft = (maskWidth / 2) - ($('#pop_box').width() / 2);
        $('#pop_box').css({ top: popTop, left: popLeft }).slideDown(600);        
        $('#back_over').css("display","block");
        title = title ==""?"免费体验网店系统后台功能":title;
    	$('#pop_box .pop_tit em').html(title);
    	$('#mudi').val(mudi);
    }
    function popclose(){
    	$('#pop_box').slideUp(600,function(){
    		$('#back_over').css("display","none");
    	})
    }
</script>
<!--弹出层结束-->
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