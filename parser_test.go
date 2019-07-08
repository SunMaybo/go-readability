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
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<title>“水氢发动机”下  地方招商引资痛点何在_中国科技网</title>
<meta name="keywords" content="招商引资,企业,招商" />
<meta name="description" content="  近日，南阳“水氢发动机”事件成为了社会关注的热点。从目前媒体披露的情况看，南阳市高新区管委会官员称相关技术是科技部“973计划”立项项目，而专家指出这个研究仅停留在理论层面；当事人强调水制氢试验车是第四代，基本各项功能都已经成熟。这个市" />
<link href="/96kaifa/css/basic.css" type="text/css" rel="stylesheet" />
<script src="/96kaifa/js/jquery.js" type="text/javascript"></script>
<script src="/96kaifa/js/common.js" type="text/javascript"></script>
<script src="/96kaifa/js/banner.js" type="text/javascript"></script>
<script src="/96kaifa/js/article.js" type="text/javascript"></script>
<script type="text/javascript">
    (function(){var ua=navigator.userAgent.toLowerCase();var bIsIpad=ua.match(/ipad/i)=="ipad";var bIsIphoneOs=ua.match(/iphone os/i)=="iphone os";var bIsAndroid=ua.match(/android/i)=="android";var bIsWM=ua.match(/windows mobile/i)=="windows mobile";if(bIsIpad||bIsIphoneOs||bIsAndroid||bIsWM){window.location.href="http://m.lztsh.com/news/cjpl/36709.html"}})();
</script>
</head>
<body>
<!--头部-->
<script src="/96kaifa/js/jquery.js" type="text/javascript"></script>
<div class="h_infoNav">
  <div class="h_masthead">
    <p><span class="cRed">搜集最新资讯！索引互联网信息！</span></p>
    <div class="h_set"> <span class="cGray">
      <div class="toplinks"><a href="/e/member/login/">登录</a>&#160;<a href="/e/member/register/index.php?groupid=1&button">注册</a>&#160; <span><a href="/shijian/" title="事件记录">事件记录</a></span></div>
    </div>
  </div>
</div>
<!--logo-->
<div class="h_logo">
  <div class="thelogo"> <a href="/" title="中国科技网"><img src="/96kaifa/images/logo.png" alt="中国科技网"/></a>
  </div>
  <div class="search">
    <form name="form1" method="post" action="/e/search/" target="_blank">
      <dl>
        <dd class="search-ico"></dd>
        <dd class="search-input">
          <input type="text" name="keyboard"  id="bdcsMain" value="请输入搜索内容" autocomplete="off" onfocus="if(this.value=='请输入搜索内容'){this.value=''}" onblur="if(this.value==''){this.value='请输入搜索内容'}">
        </dd>
      </dl>
      <a type="submit"><img src="/96kaifa/images/btn-search.gif" alt="搜索" style="width:73px;height:40px;"></a>
	  <div style="display:none">
	   <input type="hidden" value="title" name="show"></input>
	   <input type="hidden" name="tbname" value="news">
	   <input type="hidden" name="tempid" value="1">
	  </div>
    </form>
  </div>
</div>
<!--导航-->
<div id="menu">
  <ul id="nav">
    <li class="mainlevel thisclass" style="background:none"> <a href="/" title="中国科技网">首页</a> </li>
    <li class="mainlevel" id="mainlevel_01"><a href="/news/" >经济新闻</a>
      <ul id="sub_01" style="display: none;">
        <li><a href="/news/zhongguo/">国内资讯</a></li>
        <li><a href="/news/guoji/">国际资讯</a></li>
        <li><a href="/news/cjpl/">专家评论</a></li>
        <li><a href="/news/jinrong/">金融曝光</a></li>
      </ul>
    </li>
    <li class="mainlevel" id="mainlevel_02"><a href="/finance/" >财经资讯</a>
      <ul id="sub_02" style="display: none;">
        <li><a href="/finance/jrcj/"> 今日财经</a></li>
        <li><a href="/finance/cjpl/">财经评论</a></li>
        <li><a href="/finance/gszj/">股市证券</a></li>
        <li><a href="/finance/cjbt/">财经曝光</a></li>
      </ul>
    </li>
    <li class="mainlevel" id="mainlevel_03"><a href="/qiye/" >产业资讯</a>
      <ul id="sub_03" style="display: none;">
        <li><a href="/qiye/qydt/"> 企业动态</a></li>
        <li><a href="/qiye/cyjj/">产业经济</a></li>
        <li><a href="/qiye/cybg/">产业曝光</a></li>
        <li><a href="/qiye/qybg/">企业曝光</a></li>
      </ul>
    </li>
    <li class="mainlevel" id="mainlevel_03"><a href="/P2P/" >网贷资讯</a>
      <ul id="sub_03" style="display: none;">
        <li><a href="/P2P/wdxw/">网贷新闻</a></li>
        <li><a href="/P2P/ptdt/">平台动态</a></li>
        <li><a href="/P2P/zhouchaolishi/">行业观点</a></li>
        <li><a href="/P2P/minguolishi/">事件档案</a></li>
        <li><a href="/P2P/shijielishi/">考察评测</a></li>
        <li><a href="/P2P/wdbg/">网贷曝光</a></li>
      </ul>
    </li>    <li class="mainlevel" id="mainlevel_03"><a href="/hulianwang/" >IT资讯</a>
      <ul id="sub_03" style="display: none;">
        <li><a href="/hulianwang/itxw/">IT新闻</a></li>
        <li><a href="/hulianwang/itrz/">IT融资</a></li>
        <li><a href="/hulianwang/wzbg/">网站曝光</a></li>
      </ul>
    </li>
    <li class="mainlevel" id="mainlevel_05"><a href="/yule/" >娱乐新闻</a>
      <ul id="sub_05" style="display: none;">
        <li><a href="/yule/dianshi/">电视新闻</a></li>
        <li><a href="/yule/dianying/">电影新闻</a></li>
        <li><a href="/yule/zongyi/">综艺新闻</a></li>
        <li><a href="/yule/bagua/">娱乐八卦</a></li>
      </ul>
    </li>
    <li class="mainlevel" id="mainlevel_05"><a href="/jiankang/" >行业资讯</a>
      <ul id="sub_05" style="display: none;">
        <li><a href="/jiankang/jkbj/">消费资讯</a></li>
        <li><a href="/jiankang/aqls/">食品健康</a></li>
        <li><a href="/jiankang/yaopin/">医疗药品</a></li>
        <li><a href="/jiankang/ypbg/">315曝光</a></li>
    
      </ul>
    </li>
    <div class="clear"></div>
  </ul>
</div>
<!--logo-->
<div class="main">
  <div class="postion">您当前位置：<a href="/">首页</a>&nbsp;>&nbsp;<a href="/news/">经济资讯</a>&nbsp;>&nbsp;<a href="/news/cjpl/">专家评论</a> > “水氢发动机”下  地方招商引资痛点何在</div>
  <div class="wraper">
    <div class="box_left">
      <div class="art_top"><h1>“水氢发动机”下  地方招商引资痛点何在</h1></div>

      <div class="art_ext"> <span>时间：2019-06-01</span> <span>作者：佚名</span> <span>浏览：
        <script src="/e/public/ViewClick/?classid=5&id=36709&addclick=1"></script>
        </span>
      </div>
      <div class="body">
        <script src=/d/js/acmsd/thea1.js></script>
        <div class="topdes">  近日，南阳“水氢发动机”事件成为了社会关注的热点。从目前媒体披露的情况看，南阳市高新区管委会官员称相关技术是科技部“973计划”立项项目，而专家指出这个研究仅停留在理论层面；当事人强调水制氢试验车是第四代，基本各项功能都已经成熟。这个市</div>
        <p>　　近日，南阳“水氢发动机”事件成为了社会关注的热点。从目前媒体披露的情况看，南阳市高新区管委会官员称相关技术是科技部“973计划”立项项目，而专家指出这个研究仅停留在理论层面；当事人强调水制氢试验车是第四代，基本各项功能都已经成熟。这个市委书记“点赞”的“高科技产品”究竟如何还有待有关方面的进一步澄清，还公众一个真相。但是，这个事件却折射到我国经济发展的“痛点”，就是如何规范地方政府招商引资行为，防止部分官员“乱作为”，不让逐利的资本“有空可钻”。</p>
 <p>　　不可否认，地方政府招商引资是我国经济发展的一个动力。以至于张五常老先生在冥思苦想之后，发明了“县域经济竞争”理论来解释中国改革开放以来的经济成就。诚如所言，各级地方政府把当地作为一个“大企业”来运营，通过招商引资，发展产业经济，创造更多就业，增加居民的收入，为当地带来了经济繁荣，中国经济发展与地方政府主动作为、开拓创新是分不开的。这是被无数实践所反复证明了的，从近期的事例看，贵阳大数据产业发展是有力的证明，贵阳的地理位置、发展程度、人文环境在全国众多的城市中并不占优势，但由于贵阳市政府作出了前瞻性的部署，比别人深想一层，先行一步，引进了<span id="stock_NASDAQ_INTC7">英特尔</span><span id="quote_INTC7"></span>、思爱普、富士康等一批世界大型企业，构建起较为完整的大数据产业链，2018年贵阳市大数据企业主营业务收入达千亿元，2025年全市大数据产值将破万亿元，实现了跨越式发展。从更早的时期看，改革开放之初，深圳从一个小渔村发展到今天的大都市，很大程度上是依靠招商引资，引进了香港和国际的<span id="Info.385">外商直接投资</span>，通过外部的资本和内部的劳动力结合，创造了“经济奇迹”，这里面政府是功不可没的。</p>
 <p>　　招商引资是个技术活儿，并不是每个地方政府都会招商，也不是都能成功。招商引资首先要找到理想的企业，一般而言，各级地方政府都想“招大引强”，世界500强、中国500强往往是各级政府紧盯的目标。这些企业在产业链上处于核心地位，一个核心企业引进来，会带来一批上下游的中小企业入驻，从而实现做大产业、发展经济的目标。但是，好的企业数量是有限的，企业的产能也是有限的，各级政府都想引进，就存在“僧多粥少”的问题，于是，各地开始在招商引资方式上进行探索创新，出现了产业<span id="Info.3293">基金</span>招商、产业链招商、飞地招商、校友招商等多种形式。武汉是校友招商的首创地，通过在汉高校，对外地校友企业家进行招商，加之武汉的科研人才优势，取得了不俗成效，引起各地纷纷效仿。在经济发展的大潮中，招商引资也逐渐成了检验各级地方政府能力的“试金石”，招商引资是各级政府官员的“必修课”。</p>
 <p>　　凡事都有两面，兴一利必生一弊，在招商引资的过程中，也出现了一些问题。主要有两个方面：一是恶性竞争。由于各地招商引资蜂拥而至，企业可以“待价而沽”，各地只能是拼优惠政策，于是开始出现“零地价”，企业可以不付出成本就能获得土地使用权，再接着出现了“拎包入驻”，政府把企业的厂房先建好。不少地方政府在招商引资上投入了大量的资金，但最后却没有收获到预期的成果，造成了巨大的浪费。二是名不副实。一些不法商人利用地方政府招商引资的迫切心理，包装一些“高大上”的企业，占用大量的土地，然后穷尽各种办法，钻政策的空子“瞒天过海”，把土地变现从中牟取暴利，导致了国有资产和资源的大量流失。在当前经济下行的态势下，政府发展经济的动力比以往更大，在招商引资上也存在“捡到篮子的都是菜”的现象。尤其近年来，经济处于新旧动能转换的换挡期，各地都在寻找新动能，一些贴上流行的高科技标签的企业容易被各地政府“青睐”，也容易产生“鱼目混珠”的问题。</p>
 <p>　　要彻底解决这些乱象，根本在于树立正确的政绩观。招商引资的背后是官员的考核和激励机制，都是来自经济发展这个指挥棒，考核更多的是强调约束，而激励则更多的是引导。在这样的制度下，各级政府官员主要关心的是各项经济指标，企业引进来了，工厂开工了，经济指标提高了，政绩也就上去了。对于企业本身发展如何、是否符合当地实际、是否名不副实，则不是考虑的重点。针对这些问题，要切实加快改革步伐。一要进一步加强顶层设计，明确规则、制定标准、指导操作、开展评估，通过一套有效的机制来对各地的招商引资进行规范，防止招商引资偏离正常轨道。二要在政府和市场之间做好区分，让政府和市场做各自擅长的事情。要充分利用市场上的专业招商团队，把专业性的工作交给市场主体，政府主要负责宏观管理和微观督察。三要提高政府官员素质。对具体从事招商引资工作的干部进行系统培训，尤其是要加强科技知识学习，充实一批科技人才到招商队伍中来，形成招商引资“智力”支撑。</p><p class="em_media">（文章来源：中国经营报）</p>
        <!--原文标题-->
        
            
                            <br />【本文为网友上传或转载其它网站，文章版权归原作者及原出处所有。文章系作者个人观点，不代表本网站立场，转载请联系原作者及原出处获得授权。有任何疑问都请联系（125021638@qq.com）】		<div class="art_tag"> 相关标签 
<a href='/e/tags/?tagname=招商引资' target='_blank'>招商引资</a><a href='/e/tags/?tagname=企业' target='_blank'>企业</a><a href='/e/tags/?tagname=招商' target='_blank'>招商</a>		 </div>
        <div class="pages2">
          <ul>
		  	
          </ul>
        </div>
        <div class="art_tag"> </div>
        <script src=/d/js/acmsd/thea2.js></script>
        <div style="clear: both;"></div>
      </div>
      <div class="handover"> 上一篇：<a href="/news/cjpl/36708.html" target="_blank">坚决反对美方霸凌主义 重大原则问题决不让步——专家评析中美经贸热点问题</a>  下一篇：<a>没有了</a> </div>
     <div style="height:32px; width:500px;" class="tag_fl"><strong><a >相关文章</a></strong></div>
	  <div class=" tag_fl">
        <ul class="listbox">
          <li><a href="/news/cjpl/36709.html" title="“水氢发动机”下  地方招商引资痛点何在">“水氢发动机”下  地方招商引资痛点何在</a></li>
          <li><a href="/news/cjpl/36706.html" title="提升城市流动人口的社会接纳水平">提升城市流动人口的社会接纳水平</a></li>
          <li><a href="/news/cjpl/36708.html" title="坚决反对美方霸凌主义 重大原则问题决不让步——专家评析中美经贸热点问题">坚决反对美方霸凌主义 重大原则问题决不让步——专家评析中美经贸热点问题</a></li>
          <li><a href="/news/cjpl/34937.html" title="综艺营销迎来“至暗时刻” 价值王国会由短视频接管吗？">综艺营销迎来“至暗时刻” 价值王国会由短视频接管吗？</a></li>
          <li><a href="/news/cjpl/34944.html" title="交行首席经济学家连平：小型银行应坚持本地化经营导向 鼓励中小民营资本分散持股">交行首席经济学家连平：小型银行应坚持本地化经营导向 鼓励中小民营资本分散持股</a></li>
          <li><a href="/news/cjpl/34941.html" title="58贝壳争夺战：抢人、互诉与喊话带来的流量物语">58贝壳争夺战：抢人、互诉与喊话带来的流量物语</a></li>
        </ul>
      </div>
      <div class=" tag_fr">
        <ul class="listbox">
          <li><a href="/news/cjpl/34946.html" title="清华大学白重恩：金融不应过多承担财政职能">清华大学白重恩：金融不应过多承担财政职能</a></li>
          <li><a href="/news/cjpl/34947.html" title="农村网络零售额不断攀升 下沉市场用户的春天？">农村网络零售额不断攀升 下沉市场用户的春天？</a></li>
          <li><a href="/news/cjpl/34948.html" title="股市派对不会持续太久 白银大幅反弹势在必行">股市派对不会持续太久 白银大幅反弹势在必行</a></li>
          <li><a href="/news/cjpl/34949.html" title="宏观杠杆率再攀升 不能仅仅金融去杠杆">宏观杠杆率再攀升 不能仅仅金融去杠杆</a></li>
          <li><a href="/news/cjpl/34951.html" title="金融供给侧结构性改革还须防止供需错位">金融供给侧结构性改革还须防止供需错位</a></li>
          <li><a href="/news/cjpl/34950.html" title="多措并举助力企业创新">多措并举助力企业创新</a></li>
        </ul>
      </div>
      <div class="art_lbox"> <span class="art_tit"><a>热门推荐</a></span>
        <div class="art_pic">
          <ul>
            <li><a href="/qiye/cyjj/14937.html"><img src="/d/file/201711091221/201711081613132001768410.jpg" /></a>
              <p><a href="/qiye/cyjj/14937.html">汽车金融渗透率持续提升 2017年贷款规模将超万亿</a></p>
            </li>
            <li><a href="/qiye/cyjj/14951.html"><img src="/d/file/201711091222/20171108155345274005438.jpg" /></a>
              <p><a href="/qiye/cyjj/14951.html">2017年10月新能源乘用车销量6.5万辆 同比增长翻番</a></p>
            </li>
            <li><a href="/qiye/cyjj/14965.html"><img src="/d/file/201711091221/201711081543271247843453.jpg" /></a>
              <p><a href="/qiye/cyjj/14965.html">中国医药行业市场规模及市场前景分析</a></p>
            </li>
            <li><a href="/qiye/cyjj/14960.html"><img src="/d/file/201711091222/201711081545451967705376.jpg" /></a>
              <p><a href="/qiye/cyjj/14960.html">全国能源生产总量历年数据分析：2016年能源产量明显下降</a></p>
            </li>
          </ul>
        </div>
        <div class=" rel_art">
          <ul>
            <li><a href="/qiye/cyjj/14960.html" title="全国能源生产总量历年数据分析：2016年能源产量明显下降">全国能源生产总量历年数据分析：2016年能源产量明显下降</a></li>
            <li><a href="/qiye/cyjj/14966.html" title="工业应用领域优势明显 机器视觉市场将超200亿">工业应用领域优势明显 机器视觉市场将超200亿</a></li>
            <li><a href="/hulianwang/itxw/14272.html" title="Snap第三季度用户增长低于预期 盘后股价暴跌17%">Snap第三季度用户增长低于预期 盘后股价暴跌17%</a></li>
            <li><a href="/hulianwang/itxw/14271.html" title="网络现金贷仍存灰色地带 隐蔽收费名目繁多">网络现金贷仍存灰色地带 隐蔽收费名目繁多</a></li>
            <li><a href="/hulianwang/itxw/14270.html" title="股价大跌八成 网宿科技高管3个月或套现4亿">股价大跌八成 网宿科技高管3个月或套现4亿</a></li>
            <li><a href="/hulianwang/itxw/14269.html" title="现金贷数据江湖:1.5元买到一条借款人信息">现金贷数据江湖:1.5元买到一条借款人信息</a></li>
            <li><a href="/hulianwang/itxw/14267.html" title="不仅仅是中国，现在亚洲年轻人大多热衷网上购物">不仅仅是中国，现在亚洲年轻人大多热衷网上购物</a></li>
            <li><a href="/hulianwang/itxw/14266.html" title="美国科技五巨头业绩"亮眼" 对其他企业是好事吗">美国科技五巨头业绩"亮眼" 对其他企业是好事吗</a></li>
            <li><a href="/hulianwang/itxw/14265.html" title="内地比特币交易关闭后 投机者涌向香港 场外交易">内地比特币交易关闭后 投机者涌向香港 场外交易</a></li>
            <li><a href="/hulianwang/itxw/14264.html" title="百度推无人驾驶技术背后 陆奇详解AI的商业化逻">百度推无人驾驶技术背后 陆奇详解AI的商业化逻</a></li>
          </ul>
        </div>
      </div>
      <div class="project">
        <div class="tab">
          <div id="bg" class="xi1">
            <div id="font1" class="tab1" onmouseover="setTab03Syn(1);document.getElementById('bg').className='xi1'" style="color: rgb(27, 109, 157);">时事资讯</div>
            <div id="font2" class="tab2" onmouseover="setTab03Syn(2);document.getElementById('bg').className='xi2'" style="color: rgb(0, 0, 0);">财经头条</div>
            <div id="font3" class="tab3" onmouseover="setTab03Syn(3);document.getElementById('bg').className='xi3'" style="color: rgb(0, 0, 0);">网贷资讯</div>
            <div id="font4" class="tab4" onmouseover="setTab03Syn(4);document.getElementById('bg').className='xi4'" style="color: rgb(0, 0, 0);">产业经济</div>
            <div id="font5" class="tab5" onmouseover="setTab03Syn(5);document.getElementById('bg').className='xi5'" style="color: rgb(0, 0, 0);">IT科技</div>
            <div id="font6" class="tab6" onmouseover="setTab03Syn(6);document.getElementById('bg').className='xi6'" style="color: rgb(0, 0, 0);">娱乐八卦</div>
          </div>
          <div id="TabTab03Con1" style="display: block;">
            <div class="pro_txt">
                <a href="/news/guoji/14809.html" title="亚马逊业绩超预期 贝索斯成世界新首富"><img src="/d/file/201711082209/Quote.png" width="320" height="215" alt="亚马逊业绩超预期 贝索斯成世界新首富"/>
                <p>亚马逊业绩超预期 贝索斯成世界新首富</p>
                </a>
                <ul>
                  <li><a href="/news/cjpl/36709.html" title="“水氢发动机”下  地方招商引资痛点何在">“水氢发动机”下  地方招商引资痛点何在</a></li>
                  <li><a href="/news/cjpl/36706.html" title="提升城市流动人口的社会接纳水平">提升城市流动人口的社会接纳水平</a></li>
                  <li><a href="/news/cjpl/36708.html" title="坚决反对美方霸凌主义 重大原则问题决不让步——专家评析中美经贸热点问题">坚决反对美方霸凌主义 重大原则问题决不让步——专家评析中美经贸热点问题</a></li>
                  <li><a href="/news/zhongguo/34774.html" title="5月份PMI低于荣枯线？莫慌！进来看看金融业“首席们”为啥有信心">5月份PMI低于荣枯线？莫慌！进来看看金融业“首席们”为啥有信心</a></li>
                  <li><a href="/news/zhongguo/34779.html" title="30万亿元养老金入市是“10年目标” 专家称有助于建设有韧性资本市场">30万亿元养老金入市是“10年目标” 专家称有助于建设有韧性资本市场</a></li>
                  <li><a href="/news/guoji/34618.html" title="美国5月密歇根大学5年通胀预期终值 2.6">美国5月密歇根大学5年通胀预期终值 2.6</a></li>
                  <li><a href="/news/cjpl/34937.html" title="综艺营销迎来“至暗时刻” 价值王国会由短视频接管吗？">综艺营销迎来“至暗时刻” 价值王国会由短视频接管吗？</a></li>
              </ul>
            </div>
          </div>
          <div id="TabTab03Con2" style="display: none;">
            <div class="pro_txt">
                <a href="/finance/cjpl/10539.html" title="央视评论:看楼市再看天气 这次风向真的变了吗？"><img src="http://cms-bucket.nosdn.127.net/catchpic/f/f0/f0b62fc4032b6995e2ec4f2d9608301e.jpg?imageView&thumbnail=550x0" width="320" height="215" alt="央视评论:看楼市再看天气 这次风向真的变了吗？"/>
                <p>央视评论:看楼市再看天气 这次风向真的变了吗？</p>
                </a>
                <ul>
                  <li><a href="/finance/gszj/35830.html" title="商务部：中国将建立不可靠实体清单制度">商务部：中国将建立不可靠实体清单制度</a></li>
                  <li><a href="/finance/gszj/35829.html" title="北向资金多事之月 创多项记录">北向资金多事之月 创多项记录</a></li>
                  <li><a href="/finance/gszj/35828.html" title="这个题材爆发！龙头12天10涨停，暴涨1.8倍！一类股却成跌停重灾区">这个题材爆发！龙头12天10涨停，暴涨1.8倍！一类股却成跌停重灾区</a></li>
                  <li><a href="/finance/gszj/35827.html" title="烟草系杀入消费金融领域！继百度金融参股哈银消费 红塔银行也入股苏宁消费金融">烟草系杀入消费金融领域！继百度金融参股哈银消费 红塔银行也入股苏宁消费金融</a></li>
                  <li><a href="/finance/gszj/35826.html" title="上交所：本周加大信息披露和股价异常的联动监管">上交所：本周加大信息披露和股价异常的联动监管</a></li>
                  <li><a href="/finance/gszj/35825.html" title="中国证券投资基金业协会会长：建议基金产品层面取消增值税">中国证券投资基金业协会会长：建议基金产品层面取消增值税</a></li>
                  <li><a href="/finance/gszj/35824.html" title="集体重挫！全球股市“黑五月”！北上资金净卖537亿创纪录！6月会好么？">集体重挫！全球股市“黑五月”！北上资金净卖537亿创纪录！6月会好么？</a></li>
              </ul>
            </div>
          </div>
          <div id="TabTab03Con3" style="display: none;">
            <div class="pro_txt">
                <a href="/P2P/wdxw/11268.html" title="汇聚百家平台，网贷天眼双十一阶段性战报喜人"><img src="http://img.p2peye.com/2017/11/07/d333aeee153c253b2f33d916d340dc36.png" width="320" height="215" alt="汇聚百家平台，网贷天眼双十一阶段性战报喜人"/>
                <p>汇聚百家平台，网贷天眼双十一阶段性战报喜人</p>
                </a>
                <ul>
                  <li><a href="/P2P/minguolishi/36704.html" title="信而富发布5月回款公告 全面催收短期消费贷">信而富发布5月回款公告 全面催收短期消费贷</a></li>
                  <li><a href="/P2P/minguolishi/36703.html" title="近百名团贷网利益相关人员非法聚集 3人被刑拘">近百名团贷网利益相关人员非法聚集 3人被刑拘</a></li>
                  <li><a href="/P2P/minguolishi/36702.html" title="团贷网最新通报：警方已对群众线索开展核查">团贷网最新通报：警方已对群众线索开展核查</a></li>
                  <li><a href="/P2P/minguolishi/36701.html" title="警方通报“利民网”“零钱罐”案件最新进展">警方通报“利民网”“零钱罐”案件最新进展</a></li>
                  <li><a href="/P2P/minguolishi/36700.html" title="警方通报银都控股案进展：一犯罪嫌疑人被抓">警方通报银都控股案进展：一犯罪嫌疑人被抓</a></li>
                  <li><a href="/P2P/minguolishi/36699.html" title="人人贷已对接金融服务平台">人人贷已对接金融服务平台</a></li>
                  <li><a href="/P2P/shijielishi/36698.html" title="你我贷加仓">你我贷加仓</a></li>
              </ul>
            </div>
          </div>
          <div id="TabTab03Con4" style="display: none;">
            <div class="pro_txt">
                <a href="/qiye/cyjj/14965.html" title="中国医药行业市场规模及市场前景分析"><img src="/d/file/201711091221/201711081543271247843453.jpg" width="320" height="215" alt="中国医药行业市场规模及市场前景分析"/>
                <p>中国医药行业市场规模及市场前景分析</p>
                </a>
                <ul>
                  <li><a href="/qiye/cyjj/31277.html" title="不良率突飙升最高超20% 农商行打寒颤">不良率突飙升最高超20% 农商行打寒颤</a></li>
                  <li><a href="/qiye/cyjj/31280.html" title="北京扩大特殊病用药报销范围 125类药品9月起新纳入报销">北京扩大特殊病用药报销范围 125类药品9月起新纳入报销</a></li>
                  <li><a href="/qiye/cyjj/31281.html" title="深圳开出区块链电子发票">深圳开出区块链电子发票</a></li>
                  <li><a href="/qiye/cyjj/31284.html" title="百白破疫苗相关科普知识(三)：接种百白破疫苗必须要知道的几个事实">百白破疫苗相关科普知识(三)：接种百白破疫苗必须要知道的几个事实</a></li>
                  <li><a href="/qiye/cyjj/31286.html" title="百白破疫苗相关科普知识(一)：一文了解什么是百白破">百白破疫苗相关科普知识(一)：一文了解什么是百白破</a></li>
                  <li><a href="/qiye/cyjj/31287.html" title="8月10日海安兄弟锦纶长丝报价持稳">8月10日海安兄弟锦纶长丝报价持稳</a></li>
                  <li><a href="/qiye/cyjj/31289.html" title="福州钢市一周回顾（8.03-8.10）">福州钢市一周回顾（8.03-8.10）</a></li>
              </ul>
            </div>
          </div>
          <div id="TabTab03Con5" style="display: none;">
            <div class="pro_txt">
                <a href="/hulianwang/itxw/14272.html" title="Snap第三季度用户增长低于预期 盘后股价暴跌17%"><img src="http://cms-bucket.nosdn.127.net/catchpic/c/c2/c282dc907985ad6471655944cb0ca39b.png?imageView&thumbnail=550x0" width="320" height="215" alt="Snap第三季度用户增长低于预期 盘后股价暴跌17%"/>
                <p>Snap第三季度用户增长低于预期 盘后股价暴跌17%</p>
                </a>
                <ul>
                  <li><a href="/hulianwang/itxw/36327.html" title="司库立方获信天创投、海贝资本数千万元Pre-A轮">司库立方获信天创投、海贝资本数千万元Pre-A轮</a></li>
                  <li><a href="/hulianwang/itxw/36326.html" title="喜茶苏州圆融店被查封?喜茶：制作过程未发现苍">喜茶苏州圆融店被查封?喜茶：制作过程未发现苍</a></li>
                  <li><a href="/hulianwang/itxw/36325.html" title="小米成立新集团质量办公室 任命李涛为办公室主">小米成立新集团质量办公室 任命李涛为办公室主</a></li>
                  <li><a href="/hulianwang/itxw/36324.html" title="B站开启Vlog星计划，每月100万专项奖金支持">B站开启Vlog星计划，每月100万专项奖金支持</a></li>
                  <li><a href="/hulianwang/itxw/36323.html" title="H连锁酒店获华住战略投资 "连锁化"破解单体酒店">H连锁酒店获华住战略投资 "连锁化"破解单体酒店</a></li>
                  <li><a href="/hulianwang/itxw/36322.html" title="百度与东软达成战略合作 推进人工智能规模化落">百度与东软达成战略合作 推进人工智能规模化落</a></li>
                  <li><a href="/hulianwang/itxw/36321.html" title="阿里投资VMate亿级美元 称要打造印度短视频社区">阿里投资VMate亿级美元 称要打造印度短视频社区</a></li>
              </ul>
            </div>
          </div>
          <div id="TabTab03Con6" style="display: none;">
            <div class="pro_txt">
                <a href="/yule/dianshi/12735.html" title="《美味奇缘》首播收视破1 美食题材剧集传递生活真实滋味"><img src="/d/file/201711081042/20170913091951_b22ddf63cd148f841a11129ce40dbe0c_1.jpg" width="320" height="215" alt="《美味奇缘》首播收视破1 美食题材剧集传递生活真实滋味"/>
                <p>《美味奇缘》首播收视破1 美食题材剧集传递生活真实滋味</p>
                </a>
                <ul>
                  <li><a href="/yule/bagua/34120.html" title="袁惟仁被拍到与美女约会 两人互动十分亲密">袁惟仁被拍到与美女约会 两人互动十分亲密</a></li>
                  <li><a href="/yule/bagua/34119.html" title="女神“秋雅”王智结婚啦 男友居然是圏内有名的摄影师尤奕">女神“秋雅”王智结婚啦 男友居然是圏内有名的摄影师尤奕</a></li>
                  <li><a href="/yule/bagua/34118.html" title="林志玲凌晨晒出婚纱照 引发网友猜测:这是要结婚了？">林志玲凌晨晒出婚纱照 引发网友猜测:这是要结婚了？</a></li>
                  <li><a href="/yule/bagua/34117.html" title="40岁萧淑慎与小14岁男友登记结婚 双方亲友到场见证">40岁萧淑慎与小14岁男友登记结婚 双方亲友到场见证</a></li>
                  <li><a href="/yule/bagua/34116.html" title="宋仲基婚礼甜蜜宣誓内容曝光：致我高尚的新娘">宋仲基婚礼甜蜜宣誓内容曝光：致我高尚的新娘</a></li>
                  <li><a href="/yule/bagua/34115.html" title="言承旭与林志玲复合 昔日绯闻女友晒照发文">言承旭与林志玲复合 昔日绯闻女友晒照发文</a></li>
                  <li><a href="/yule/bagua/34114.html" title="言承旭承认与林志玲复合 获林志玲妈妈点头认可">言承旭承认与林志玲复合 获林志玲妈妈点头认可</a></li>
              </ul>
            </div>
          </div>
        </div>
      </div>
      <div class="art_lbox">
        <div class="hr10"></div>
        <span class="art_tit"><a>猜你喜欢</a></span>
        <div class=" ztbox">
          <div class="zt_pic left">
            <ul>
              <li class="aniTit"><a href="/P2P/minguolishi/24687.html"><img src="http://imgmywdzj.wdzj.com/group1/M00/44/89/rBAEX1tHFsuAfz8JAAEpaMTsi1g290.jpg"/></a><span class="title-alpha"><a href="/P2P/minguolishi/24687.html">石头理财没有按时回款？有没有朋友知道是怎么回事的？</a></span></li>
              <li class="aniTit"><a href="/yule/bagua/34056.html"><img src="http://img.mingxing.com/upload/thumb/1059/317585.jpg"/></a><span class="title-alpha"><a href="/yule/bagua/34056.html">陈晓东自爆被阔太包养三十多年 主持人都惊呆</a></span></li>
            </ul>
          </div>
          <div class="ztlist left">
            <ul>
              <li><a href="/news/zhongguo/31426.html">易纲：中国将进一步推动改革扩大开放 金融部门内部控杠杆取得阶段性成效</a></li>
              <li><a href="/yule/dianying/11617.html">大表姐詹妮弗·劳伦斯指责飓风与特朗普有关 震惊美国网友</a></li>
              <li><a href="/hulianwang/itrz/25682.html">何晓飞首次披露飞步自动驾驶方案：无人货运起步 自研AI芯片</a></li>
              <li><a href="/P2P/ptdt/24752.html">用科技力量践行社会公益  凡普金科集团2013-2017企业社会责任报告亮相</a></li>
              <li class="line"></li>
              <li><a href="/finance/cjpl/10256.html">证券日报：明年退市制度或有新动向新突破</a></li>
              <li><a href="/yule/zongyi/15987.html">一周的偶像SJ挑战最新偶像潮流</a></li>
              <li><a href="/finance/jrcj/33907.html">安信信托大换血 国之杰派驻多面手</a></li>
              <li><a href="/news/zhongguo/31643.html">郑州“学区房”将何去何从</a></li>
            </ul>
          </div>
          <div class="zt_pic right">
            <ul>
              <li class="aniTit"><a href="/finance/gszj/35762.html"><img src="https://i.ssimg.cn/sscms/2019/05/30/63dca9762b62443696c40c0550bff3a5.png"/></a><span class="title-alpha"><a href="/finance/gszj/35762.html">两市缩量盘整 近300只个股日成交不足千万元</a></span></li>
              <li class="aniTit"><a href="/news/guoji/32100.html"><img src="http://www.lztsh.com/d/file/20190524/Quote.png"/></a><span class="title-alpha"><a href="/news/guoji/32100.html">咖喱味儿的40000点历史新高！印度凭什么走出十年大牛市？</a></span></li>
            </ul>
          </div>
        </div>
      </div>
      <div class="pic_pre">
        <script src="/jsfile/du114.js" type="text/javascript"></script>
      </div>
      <style>
.tj-box{display:inline-block;width:99.5%;background:#f0f0f0;border-radius: 5px;border:1px solid #e1e1e1;}
.tj-title{float:left;width:5.4%;background:#FFF;padding:8px 6px 8px 8px;border-radius: 5px;border-right:1px solid #e1e1e1;}
.tj-title strong{ float:left;font-weight: normal;font-size: 14px;
color: #1b6d9d;letter-spacing: 4px;line-height: 25px;}
.tj-txt{float:right;width:91.6%; height:58px; overflow:hidden;padding:3px 2px 3px 3px;}
.tj-txt a{display:inline-block;font-size:12px;color:#333;padding:2px;text-decoration:none;}
.tj-txt a:hover{ color:#1b6d9d;}
</style>
    </div>
    <div class="box_right">
      <script>
			article_zhuanqu();
	</script>
      <div class="read"> <span class="mark">热门标签</span>
        <div class="tag1">
		<a target="_blank" href="/e/tags/?tagname=[db:关键词]" title="[db:关键词]">[db:关键词]</a>
		<a target="_blank" href="/e/tags/?tagname=中国" title="中国">中国</a>
		<a target="_blank" href="/e/tags/?tagname=平台" title="平台">平台</a>
		<a target="_blank" href="/e/tags/?tagname=亿元" title="亿元">亿元</a>
		<a target="_blank" href="/e/tags/?tagname=互联网" title="互联网">互联网</a>
		<a target="_blank" href="/e/tags/?tagname=信息" title="信息">信息</a>
		<a target="_blank" href="/e/tags/?tagname=来源" title="来源">来源</a>
		<a target="_blank" href="/e/tags/?tagname=给我们" title="给我们">给我们</a>
		<a target="_blank" href="/e/tags/?tagname=金融" title="金融">金融</a>
		<a target="_blank" href="/e/tags/?tagname=美国" title="美国">美国</a>
		<a target="_blank" href="/e/tags/?tagname=银行" title="银行">银行</a>
		<a target="_blank" href="/e/tags/?tagname=融资" title="融资">融资</a>
		<a target="_blank" href="/e/tags/?tagname=天眼" title="天眼">天眼</a>
		<a target="_blank" href="/e/tags/?tagname=首页" title="首页">首页</a>
		<a target="_blank" href="/e/tags/?tagname=企业" title="企业">企业</a>
		<a target="_blank" href="/e/tags/?tagname=未找到" title="未找到">未找到</a>
		<a target="_blank" href="/e/tags/?tagname=此页面" title="此页面">此页面</a>
		<a target="_blank" href="/e/tags/?tagname=投资" title="投资">投资</a>
		<a target="_blank" href="/e/tags/?tagname=编辑" title="编辑">编辑</a>
		<a target="_blank" href="/e/tags/?tagname=骗子" title="骗子">骗子</a>
		</div>
      </div>
      <div class="read">
        <script src=/d/js/acmsd/thea3.js></script>
      </div>
      <div class="read"> <span class="mark">阅读排行</span>
        <div class="tab">
          <div id="TabTab03Con1" style="width:320px;">
            <ul class="tab_txt">
              <li class="one"><em>1</em> <a href="/hulianwang/itxw/21686.html" title="Pokémon Go开发商获2亿美元融资 加入独角兽俱">Pokémon Go开发商获2亿美元融资 加</a> <span><a href="/hulianwang/itxw/21686.html"><img src="http://cms-bucket.nosdn.127.net/c970194e2d084b3c88698e6bdf4e775420171125141334.jpeg?imageView thumbnail=550x0" width="100" height="71" alt="Pokémon Go开发商获2亿美元融资 加入独角兽俱"/></a>
                <p>网易科技讯11月25日消息，据金融时报报道，PokémonGo游戏...<a class="deta" href="/hulianwang/itxw/21686.html" >详细>></a></p>
                </span> </li>
              <li class="one"> <em>2</em> <a href="/jiankang/yaopin/19606.html" title="天演药业：独辟蹊径，创新生物药赛道的“中国样本”">天演药业：独辟蹊径，创新生物药赛道的</a> </li>
              <li class="one"> <em>3</em> <a href="/P2P/shijielishi/10938.html" title="东银金服（小微金融）的爹东方银座集团具体有多少资产和负债？">东银金服（小微金融）的爹东方银座集团</a> </li>
              <li> <em>4</em> <a href="/P2P/wdxw/9680.html" title="这5种食物，不但可以防癌抗癌，而且非常神奇，可惜很少人知道">这5种食物，不但可以防癌抗癌，而且非常</a> </li>
              <li> <em>5</em> <a href="/P2P/wdxw/11267.html" title="美IPO市场今年最忙一周：搜狗等11家公司将上市">美IPO市场今年最忙一周：搜狗等11家公</a> </li>
              <li> <em>6</em> <a href="/jiankang/jkbj/19893.html" title="和海鸥一起奔跑！中国第一个全程海景马拉松成功举办">和海鸥一起奔跑！中国第一个全程海景</a> </li>
              <li> <em>7</em> <a href="/P2P/wdxw/11268.html" title="汇聚百家平台，网贷天眼双十一阶段性战报喜人">汇聚百家平台，网贷天眼双十一阶段性</a> </li>
              <li> <em>8</em> <a href="/news/guoji/14343.html" title="博通1300亿美元“豪娶”高通 科技行业最大并购能否成功？">博通1300亿美元“豪娶”高通 科技行</a> </li>
              <li> <em>9</em> <a href="/P2P/wdxw/11266.html" title="招联、马上、玖富、苏宁等参战“虚拟信用卡”">招联、马上、玖富、苏宁等参战“虚</a> </li>
              <li> <em>10</em> <a href="/jiankang/jkbj/19899.html" title="欧玛士电器：极简为傲&nbsp;差异化精工单品系列">欧玛士电器：极简为傲 差异化精工单品</a> </li>
            </ul>
          </div>
        </div>
      </div>
      <div class="pic">
        <div class="read">
          <script src=/d/js/acmsd/thea4.js></script>
        </div>
        <div class="read"> <span class="mark">图片推荐</span>
          <div class="read">
            <div class="pic_rec">
              <ul>
          <li><a href="/qiye/cyjj/14965.html" title="中国医药行业市场规模及市场前景分析"><img src="/d/file/201711091221/201711081543271247843453.jpg" width="140" height="99" alt="中国医药行业市场规模及市场前景分析"/></a>
            <p>  医药制造业是指原料</p>
          </li>
          <li><a href="/qiye/cyjj/14966.html" title="工业应用领域优势明显 机器视觉市场将超200亿"><img src="/d/file/201711091223/201711081515441167723915.jpg" width="140" height="99" alt="工业应用领域优势明显 机器视觉市场将超200亿"/></a>
            <p>  机器视觉可以在各种</p>
          </li>
          <li><a href="/hulianwang/itxw/14272.html" title="Snap第三季度用户增长低于预期 盘后股价暴跌17%"><img src="http://cms-bucket.nosdn.127.net/catchpic/c/c2/c282dc907985ad6471655944cb0ca39b.png?imageView&thumbnail=550x0" width="140" height="99" alt="Snap第三季度用户增长低于预期 盘后股价暴跌17%"/></a>
            <p>网易科技讯11月8日消</p>
          </li>
          <li><a href="/hulianwang/itxw/14271.html" title="网络现金贷仍存灰色地带 隐蔽收费名目繁多"><img src="http://img1.cache.netease.com/cnews/css13/img/end_tech.png" width="140" height="99" alt="网络现金贷仍存灰色地带 隐蔽收费名目繁多"/></a>
            <p>（原标题：　网络“现金贷</p>
          </li>
          <li><a href="/hulianwang/itxw/14270.html" title="股价大跌八成 网宿科技高管3个月或套现4亿"><img src="http://cms-bucket.nosdn.127.net/catchpic/b/bb/bbd8389cb6e74f2b3f510f30a47b9663.jpg?imageView&thumbnail=550x0" width="140" height="99" alt="股价大跌八成 网宿科技高管3个月或套现4亿"/></a>
            <p>（原标题：股价大跌八成机</p>
          </li>
          <li><a href="/P2P/wdxw/11268.html" title="汇聚百家平台，网贷天眼双十一阶段性战报喜人"><img src="http://img.p2peye.com/2017/11/07/d333aeee153c253b2f33d916d340dc36.png" width="140" height="99" alt="汇聚百家平台，网贷天眼双十一阶段性战报喜人"/></a>
            <p>一年一度的双十一狂欢</p>
          </li>
          <li><a href="/P2P/wdxw/11267.html" title="美IPO市场今年最忙一周：搜狗等11家公司将上市"><img src="http://img.p2peye.com/2017/11/07/97ed421b68bde82ee35615cc65c9f349.jpg" width="140" height="99" alt="美IPO市场今年最忙一周：搜狗等11家公司将上市"/></a>
            <p>北京时间11月7日凌晨</p>
          </li>
          <li><a href="/P2P/wdxw/11266.html" title="招联、马上、玖富、苏宁等参战“虚拟信用卡”"><img src="http://img.p2peye.com/2017/11/07/e7710b481320135694f4df0031136f29.jpg" width="140" height="99" alt="招联、马上、玖富、苏宁等参战“虚拟信用卡”"/></a>
            <p>10月底，马上消费金融旗</p>
          </li>
          <li><a href="/news/guoji/14809.html" title="亚马逊业绩超预期 贝索斯成世界新首富"><img src="/d/file/201711082209/Quote.png" width="140" height="99" alt="亚马逊业绩超预期 贝索斯成世界新首富"/></a>
            <p>  美国电子商务巨头亚</p>
          </li>
          <li><a href="/news/guoji/14808.html" title="全球呈“双低型”增长 货币政策分化加剧"><img src="/d/file/201711082209/20171031064136650702594.jpg" width="140" height="99" alt="全球呈“双低型”增长 货币政策分化加剧"/></a>
            <p> 【全球呈“双低型”</p>
          </li>
              </ul>
            </div>
          </div>
          <span class="mark">随机推荐</span>
          <div id="TabTab03Con1" style="width:320px;">
            <ul class="tab_txt">
              <li class="one"><em>1</em><a href="/hulianwang/itxw/26856.html" title="京东商城开始实施轮值CEO制度 徐雷任首任轮值CE">京东商城开始实施轮值CEO制度 徐雷</a></li>
              <li class="one"><em>2</em><a href="/news/guoji/34643.html" title="内蒙古：农产品将“持证上岗”">内蒙古：农产品将“持证上岗”</a></li>
              <li class="one"><em>3</em><a href="/finance/gszj/30932.html" title="135亿净资金全面流出 五大板块遭遇杀跌">135亿净资金全面流出 五大板块遭遇</a></li>
              <li class="one"><em>4</em><a href="/yule/bagua/19350.html" title="周冬雨恋情疑似曝光 获小鲜肉探班两人嬉笑打闹互动不断">周冬雨恋情疑似曝光 获小鲜肉探班两</a></li>
              <li class="one"><em>5</em><a href="/yule/dianshi/15890.html" title="高庚杓曹在显确定出演tvN新月火剧Cross男主">高庚杓曹在显确定出演tvN新月火剧Cr</a></li>
              <li class="one"><em>6</em><a href="/P2P/zhouchaolishi/11311.html" title="网贷信披新标准能否升级为“国标”">网贷信披新标准能否升级为“国标”</a></li>
              <li class="one"><em>7</em><a href="/finance/gszj/28094.html" title="股市资金供需正寻求新的均衡">股市资金供需正寻求新的均衡</a></li>
              <li class="one"><em>8</em><a href="/news/cjpl/29735.html" title="我国金融市场的主要缺陷与地方政府金融的发展">我国金融市场的主要缺陷与地方政府</a></li>
              <li class="one"><em>9</em><a href="/finance/gszj/29644.html" title="上万亿元黑臭水体治理市场 地方政府该怎么投">上万亿元黑臭水体治理市场 地方政府</a></li>
              <li class="one"><em>10</em><a href="/finance/gszj/33041.html" title="收评：成交缩量沪指微涨0.02% 稀土等概念股大幅回调">收评：成交缩量沪指微涨0.02% 稀土等</a></li>
            </ul>
          </div>
        </div>
        <div class="read">
          <script src=/d/js/acmsd/thea5.js></script>
        </div>
      </div>
    </div>
  </div>
</div>
<div class="clear"></div>
<div id="j_gotop" style="display: block;"><a href="#" class="btn-top"></a></div>
<div class="clear"></div>
<div class="foot">
  <div class="foot_m">
    <p>Copyright @ 2017-2018 &nbsp;中国科技网 &nbsp;All Rights Reserved. &nbsp;www.lztsh.com  &nbsp;QQ：125021638 <script type="text/javascript">var cnzz_protocol = (("https:" == document.location.protocol) ? " https://" : " http://");document.write(unescape("%3Cspan id='cnzz_stat_icon_1269421587'%3E%3C/span%3E%3Cscript src='" + cnzz_protocol + "s22.cnzz.com/z_stat.php%3Fid%3D1269421587%26show%3Dpic' type='text/javascript'%3E%3C/script%3E"));</script>    </p>
    <p>免责声明：本站所有信息均搜集自互联网，并不代表本站观点，本站不对其真实合法性负责。如有信息侵犯了您的权益，请告知，本站将立刻删除。</p>
    <p><a href="http://www.lztsh.com" style="color:#1b6d9d; text-decoration:none;"><strong>中国科技网</strong></a>：一直以来坚持文明办网，传播健康、娱乐、感知照片！每一套作品都经过我们专业编辑的精挑细选。&nbsp;</p>
    <p><a style="display:inline-block;text-decoration:none;height:20px;line-height:20px;"><img src="/96kaifa/images/备案图标.png" style="float:left;"/><b style="float:left;height:20px;line-height:20px;margin: 0px 0px 0px 5px; color:#939393;font-weight:normal;">备案号：粤ICP备17024501号-2</b></a> </p>
  </div>
</div>
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