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
<title>张昭卸任乐创文娱CEO 孙宏斌长子孙喆一接棒_蜘蛛资讯网</title>
<meta name="keywords" content="新报跑狗图今期跑狗图,香港正版赌经侠诗,香港创富合数单双,四海图库彩色印刷区" />
<meta name="description" content="淮安蜘蛛资讯网为广大玩家提供最新、最全、最具特色的淮安蜘蛛资讯，同时还有各种八卦奇闻趣事。看蜘蛛资讯，就来淮安蜘蛛资讯网！" />
<link href="http://img.alizhizhuchi.com/templates/moban31/css/index.css" type="text/css" rel="stylesheet" />
<style>#_con{line-height: 30px;font-size:14px;}#_con img{max-width:100%;}</style>
</head>
<body>
<div class="top_heard">
    <div class="tip"><img src="/static/logo.jpg" alt="蜘蛛资讯网" title="蜘蛛资讯网" /></div>
	<div class="top_right"></div>
</div>
<div class="menu">
  <ul>
    <li><a href="/">首页</a></li>

      <LI><a href="http://www.mojbuy.com/9ih/" title="国内" >国内</a></LI>

      <LI><a href="http://qklvpks.bkhmx.com/ehh7s69p/" title="国际" >国际</a></LI>

      <LI><a href="http://pjwuoq87.gctcd.com/qhbgjdtc5/" title="军事" >军事</a></LI>

      <LI><a href="http://www.mojbuy.com/5dp/" title="财经" >财经</a></LI>

      <LI><a href="http://bn23igw5f.tbcdx.com/1037ri/" title="房产" >房产</a></LI>

      <LI><a href="http://siy.xtyzg.com/gei9w/" title="汽车" >汽车</a></LI>

      <LI><a href="http://www.mojbuy.com/iewcdy/" title="体育" >体育</a></LI>

      <LI><a href="http://9may7zy7u.zfyfh.com/ic7dyktv/" title="娱乐" >娱乐</a></LI>

      <LI><a href="http://9x49wsya.tbxwz.com/pu2k2xr6/" title="教育" >教育</a></LI>

      <LI><a href="http://www.mojbuy.com/20mek7/" title="科技" >科技</a></LI>

      <LI><a href="http://u3k00co.xtxxm.com/9hvjz3/" title="社会" >社会</a></LI>
    
  </ul>
</div>
<div class="menu_active">
    <div>蜘蛛资讯网最近更新：<a href="http://5xz5.ymyyk.com/9ccahjyuy/15933-33653-54336.html">国家卫生委员会：西藏自治区人均预期寿命提高到70.6岁</a>&nbsp;&nbsp;
        <a href="http://qznei9cd.zfjcx.com/6k0/34735-57095-51037.html">iOS 13 设计指南：深色模式篇</a>&nbsp;&nbsp;
        <a href="http://www.mojbuy.com/vyp/45155-56218-23674.html">06月18日上海金交所iAu99.99价格298.90</a>&nbsp;&nbsp;
        <a href="http://z69.kjey.top/28nieue57/20402-62764-70890.html">经济日报：放心吧！猪肉供应得到充分保证</a>&nbsp;&nbsp;
        <a href="http://sdc8qenbv.kjmi.top/m1trpxsug/9782-22389-43515.html">工业互联网红利：3200多万企业将拥有自己的智能应用程序</a>&nbsp;&nbsp;
    </div>
</div>

<div id="list">
  <div class="list_l">
    <div class="zhaopin fram">
        <div class="tit">
            <h2><a href="http://c8xm0hn.kjuf.top/8pav/">蜘蛛资讯网最新文章</a></h2>
        </div>
        <div class="zhaopin_list cont">
            <a href="http://oet.kjed.top/i106f/26902-52622-73619.html" title="为什么不推荐用勺子吃西瓜？原因在于你的摄入量。">&middot;为什么不推荐用勺子吃西瓜？原因在于你的摄入量。</a>
            <a href="http://b0gl0.kjop.top/t816/12112-16174-73529.html" title="俄罗斯媒体：四名在美国被捕的“反政变人士”因害怕面临长达一年的监禁而获释。">&middot;俄罗斯媒体：四名在美国被捕的“反政变人士”因害怕面临长达一年的监禁而获释。</a>
            <a href="http://hhq.kjknx.top/ge22/1015-74175-12569.html" title="8.9不同的电影，现实生活改编：12岁的男孩起诉他的父母，“因为他们生了我。”">&middot;8.9不同的电影，现实生活改编：12岁的男孩起诉他的父母，“因为他们生了我。”</a>
            <a href="http://www.mojbuy.com/hf59dr/44175-55459-12076.html" title="应王毅邀请 荷兰外交大臣布洛克将访华">&middot;应王毅邀请 荷兰外交大臣布洛克将访华</a>
            <a href="http://oykhk2.lhccq.top/n5jnz152c/38451-65222-68395.html" title="原高管被查+不良高企 富滇银行风控受考">&middot;原高管被查+不良高企 富滇银行风控受考</a>
            <a href="http://m0hy6e.kjjpw.top/yei/56309-62679-42993.html" title="“崩盘危机”将重现？外界担忧这个国家成目标">&middot;“崩盘危机”将重现？外界担忧这个国家成目标</a>
            <a href="http://www.mojbuy.com/105cpucn/58101-70564-84999.html" title="重磅！大湾区可享受颁布的《万元年薪政府补助30万税》规定。">&middot;重磅！大湾区可享受颁布的《万元年薪政府补助30万税》规定。</a>
            <a href="http://8ck9.kjwo.top/npaj/2650-46044-45951.html" title="海滨精品酒店荣获2018年第二届GXA-良好体验奖">&middot;海滨精品酒店荣获2018年第二届GXA-良好体验奖</a>
            <a href="http://yizsgo.kjjln.top/hdeyy82j9/33451-40008-10633.html" title="不是个例！深圳高空坠落，砸中小男孩——这样的悲剧，该如何避免">&middot;不是个例！深圳高空坠落，砸中小男孩——这样的悲剧，该如何避免</a>
            <a href="http://www.mojbuy.com/kfauaa/38690-58814-39405.html" title="山东证监局原局长徐铁被查 曾任发行监管部副主任">&middot;山东证监局原局长徐铁被查 曾任发行监管部副主任</a>
            <a href="http://1lzkyc.kjah.top/83n9/24320-29523-98288.html" title="中国证监会除尘处罚*ST中安">&middot;中国证监会除尘处罚*ST中安</a>
            <a href="http://lmlqr9i.kjcnt.top/htc/36983-71379-72767.html" title="鹏博士(600804.SH)拟向讯通联盈增资3亿元">&middot;鹏博士(600804.SH)拟向讯通联盈增资3亿元</a>
        </div>
    </div>
    <div class="clear"></div>
    <div class="zhaopin fram">
        <div class="tit">
            <h2><a href="http://6vsdzm.kjgfy.top/9i8x5drk/">蜘蛛资讯网热门文章</a></h2></div>
        <div class="xinwen_list cont">
            <ul>
                <li><a href="http://8yn1o.lhcgh.top/lmz6c6wyt/898-36110-39664.html" title="短池世锦赛4x100混傅园慧率队摘银 徐嘉余200仰仅第六_体育" >短池世锦赛4x100混傅园慧率队摘银 徐嘉余200仰仅第六_体育</a><span>06-24</span></li>
                <li><a href="http://x9ca.lhcjk.top/yt1/22891-57082-55199.html" title="对退市的强力监管触发了ST行业的风格转变" >对退市的强力监管触发了ST行业的风格转变</a><span>06-26</span></li>
                <li><a href="http://www.mojbuy.com/c78v/35883-54686-80013.html" title="中国政企数字化联盟成立大会暨北京部委央企及大型企业CIO大会即将召开" >中国政企数字化联盟成立大会暨北京部委央企及大型企业CIO大会即将召开</a><span>06-27</span></li>
                <li><a href="http://www.mojbuy.com/sth665j93/20859-51441-52352.html" title="卷取卷取卷几何" >卷取卷取卷几何</a><span>06-24</span></li>
                <li><a href="http://5dqhqpjvt.lhcjy.top/3cp3w7/5489-49011-27713.html" title="真正的大盗车" >真正的大盗车</a><span>06-25</span></li>
                <li><a href="http://ik9.lhcnf.top/rkflhp/7235-18639-52735.html" title="DNF Grunt Stew策略摘要综合公式/可用奖励列表" >DNF Grunt Stew策略摘要综合公式/可用奖励列表</a><span>06-24</span></li>
                <li><a href="http://qzruy.lhcqm.top/nwb6s/11161-65602-84132.html" title="公告选择：长营精密布局氢燃料电池金属极板江苏索普重组失败" >公告选择：长营精密布局氢燃料电池金属极板江苏索普重组失败</a><span>06-30</span></li>
                <li><a href="http://s1c.mhbhy.top/30io3/37047-71454-44520.html" title="科创板来袭 新三板路在何方？" >科创板来袭 新三板路在何方？</a><span>06-24</span></li>
                <li><a href="http://k5kspxat6.mhdkc.top/skicax/23474-33508-87969.html" title="美国股市暴跌，美国政界人士正将美国经济推入“衰退”的深渊。" >美国股市暴跌，美国政界人士正将美国经济推入“衰退”的深渊。</a><span>06-30</span></li>
                <li><a href="http://vtp.mhdpm.top/jejy9zn/28226-64193-54927.html" title="乔新迈的海军被捕了？娱乐圈浪漫炒作的运作是666。" >乔新迈的海军被捕了？娱乐圈浪漫炒作的运作是666。</a><span>06-29</span></li>
                <li><a href="http://pbn.mheg.top/e5i5aqlz/68688-73244-53784.html" title="超3亿“黑财”被查封冻结扣押" >超3亿“黑财”被查封冻结扣押</a><span>06-30</span></li>
            </ul>
        </div>
    </div>
  </div>
  <div class="list_r fram">
    <div class="tit">
      <p>您当前的位置：<a href="http://8lkvp.mhfgd.top/q8fh3ztmx/">主页</a> > <a href="http://t89m1.mhic.top/46qy9/">蜘蛛资讯网国内</a> >
      <p>
    </div>
    <div class="body cont">
      <div class="body_tit">
        <p>作者：建王密建  来源：原创   发布日期：06-29</p>
      </div>
		<div class="con">
			<div id="_con"><h1 style="text-align:center">www00676金光佛论坛</h1><h2>张昭卸任乐创文娱CEO 孙宏斌长子孙喆一接棒</h2><p><p>&nbsp;&nbsp;&nbsp;&nbsp;<div><p>　　张昭卸任乐创文娱CEO，孙宏斌长子孙喆一接棒 | 钛快讯</p><p>　　钛媒体快讯 | 6月24日消息：乐创文娱官方微博发布了原董事长、CEO张昭的辞任公告，宣布张昭因个人原因卸任乐创文娱相关职位。融创文化集团总裁孙喆一将接任乐创文CEO一职，而孙喆一正是融创中国董事长孙宏斌的长子。目前，张昭与孙喆一已完成相关工作交接。 </p><p>　　资料显示，张昭于2003年加盟<span id="stock_sz300251">光线传媒</span><span id="quote_sz300251"></span>，2006年创立光线影业，2011年离开光线影业，创立乐视影业，任职CEO及执行董事。乐视影业先后拥有过张艺谋、徐克、陆川、郭敬明等多位签约导演。2014年，乐视影业获得B轮投资，投后估值48亿元，投资方包括孙俪、邓超、黄晓明、冯绍峰等多位明星。</p><p>　　乐视影业曾计划以98亿元估值注入<span id="stock_sz300104">乐视网</span><span id="quote_sz300104"></span>（维权）（1.690， 0.00， 0.00%），但随着乐视系的溃败，这一计划也以失败告终。2018年3月27日，原乐视影业董事长、CEO张昭发表内部信，宣布公司将更名为乐创文娱，融创占股40.75%，为第一大股东。而贾跃亭旗下乐视控股持有的16.3592%乐创文娱股份，将在未来逐步通过转让或拍卖方式剥离。此后，通过司法拍卖，乐视控股持有的乐创文娱股份也悉数卖出，乐创文娱与乐视系彻底剥离。</p><p>　　由于此前的视风波，一直由乐创文娱运作的《熊出没》，其第五部并没有让乐创文娱操作。而第六部《熊出没·原始时代》再次回到乐创文娱，其在2019年春节档收获7.14亿票房，这被认为是乐创文娱重回主流视野的重要作品。</p><p>　　而近期，乐创文娱出品的影片则是6月14日上映的《秦明·生死语者》，但这部影片并不成功，票房仅为2908.7万。</p><p>　　融创文化集团成立于2018年12月。2019年2月21日，融创中国官方发布消息，融创文化集团总裁一职由融创中国董事长孙宏斌的长子孙喆一担任。</p><p>　　东方影都融创影视产业园、乐创文娱与乐融致新是融创文化集团的三块核心业务，其中乐融致新的主要业务是原来的乐视电视。此前融创文化集团表示还将增设的新内容板块，共同完成文化集团的“内容+平台+实景”战略。</p><p>　　6月，融创文化集团再添新成员。在<span id="stock_sh601595">上海电影</span><span id="quote_sh601595"></span>节期间，融创文化集团宣布投资电影《解放了》《刺杀小说家》，全新内容厂牌“融创影视”就此浮出水面。</p><p>　　《刺杀小说家》的出品方信息显示，“融创影视”的实体公司为融创未来文化娱乐（北京）有限公司，成立于2019年4月10日，法定代表人为融创中国执行董事兼行政总裁汪孟德。</p><p>　　 ci sha xiao shuo jia de chu pin fang xin xi xian shi," rong chuang ying shi" de shi ti gong si wei rong chuang wei lai wen hua yu le bei jing you xian gong si, cheng li yu 2019 nian 4 yue 10 ri, fa ding dai biao ren wei rong chuang zhong guo zhi xing dong shi jian xing zheng zong cai wang meng de.</p><p>　　在融创文化集团官方微信的介绍中，“融创影视”被放在青岛东方影都融创影视产业园、乐融致新、乐创文娱等业务板块之前，可见对这一内容厂牌的重视。</p><p>　　此外，有媒体表示，卸任后的张昭已有新的创业计划，或为复星集团投资的一家影视公司。而在电影领域，复星集团比较大的动作是此前对博纳影视和好莱坞的Studio 8均有投资。</p><p>责任编辑：陈志杰 </p></div></p></p><p>当前文章：http://0q6f7dw.mojbuy.com/lko/61857-67117-17906.html</p><p>发布时间：07:47:17</p><p><a href='http://www.www-4227c.com' target='_blank'>宝马论坛管家婆四肖期期准中特网</a>&nbsp;&nbsp;<a href='http://www.bbkgolf.com' target='_blank'>www85777王中王资料</a>&nbsp;&nbsp;<a href='http://www.www200605.com' target='_blank'>www.950999.com</a>&nbsp;&nbsp;<a href='http://www.www-4394b.com' target='_blank'>静心阁高手论坛</a>&nbsp;&nbsp;<a href='http://www.www-281114.com' target='_blank'>香港马会资料大全</a>&nbsp;&nbsp;<a href='http://www.kj717.net' target='_blank'>www.kj717.com</a>&nbsp;&nbsp;<a href='http://www.www-614411.com' target='_blank'>王中王开奖记录</a>&nbsp;&nbsp;<a href='http://www.pt606.com' target='_blank'>香港白小姐一码中特料</a>&nbsp;&nbsp;<a href='http://www.cjgdj.com' target='_blank'>www.130355.com</a>&nbsp;&nbsp;<a href='http://www.www56012a.com' target='_blank'>港最快开奖现场2019结果</a>&nbsp;&nbsp; <br/></p><p><p style='font-size:16px;'>{相关文章}</p><h3>新时代文明实践站里的大忙人</h3><p><p>&nbsp;&nbsp;&nbsp;&nbsp;<p style="text-indent:2em;">见到李燕的时候，她正在北京市延庆区香水园街道恒安社区新时代文明实践站的活动大厅里组织社区的文艺队伍进行舞蹈排练。在人群<a href='http://g9j2rfcx7.mhmi.top/pmbemn061/'><em>六和彩生肖图_蜘蛛资讯网</em></a>中，李燕很显眼，一袭黑色的衣装，衬托她的皮肤更显白皙，声音洪亮，乌黑的短发让她看起来精神干练。今年是她在恒安社区工作的第五年，也是作为新时代文明实践站文明指导员的第二个年头。</p><p style="text-indent:2em;">2018年7月，恒安社区新时代文明实践站挂牌成立，同时，恒安社区新时代文明实践站四楼的一个面积80余平米冬奥世园展览馆也正式对外开放，李燕担起了新时代文明实践文明指导员的职责。2018年12月，四楼的世园冬奥展览馆挂牌新时代文明实践基地。李燕肩上的责任更重了，她不仅要负责基地平台的日常维护发布活动，同时还要接待参观人员。</p><p style="text-indent:2em;">“自从加入了‘点单派单’系统，我们展览馆接到了更多区内外的机构或兄弟单位的参观学习点单，在给他们的讲解过程中，对我个人来说其实也是一种进步和提升。”李燕说。</p><p style="text-indent:2em;">最多的时候，李燕一天接待过三组参观团队，一天讲解下来，她的嗓子嘶哑的说不出话，让她印象最深的一次是接待一个河北宣化的参观团，看了延庆区内各地的美丽景色，参观团中一名代表羡慕的对李燕说：“你们这儿的环境太好了，在延庆生活一定很幸福吧？”李燕先是一怔，随后心中不由得升起阵阵自豪。</p><p style="text-indent:2em;">除了日常的接待讲解外，新时代文明实践站的日常工作她也搞得有声有色。恒安社区新时代文明实践站成立以后，社区对实践站的组织机构进行了优化调整。下设政策宣传、文化体育、党课宣讲、安全教育、法律咨询等七项活动内容。李燕主要负责活动的组织和开展，以及<a href='http://yjib.mhvq.top/xjq/'><em>京港印刷图源_蜘蛛资讯网</em></a>了解民情民意根据实际情况在志愿服务平台进行点单，邀请区级各类志愿者为社区开展活动提供便利条件。</p><p style="text-indent:2em;">恒安社区里有一支舞蹈队，以前大家排练都是在广场或者随便挑个人少的地方就进行了，再加上没有专业的舞蹈老师进行指导，除了几个特别的舞蹈爱好者以外，队员流失严重。李燕得知了这一情况，主动找到舞蹈队的负责人宋芳珍，拉着她们来新时代文明实践基地进行<a href='http://qtbub.tbxwz.com/mz694f/'><em>大森林心水论谈_蜘蛛资讯网</em></a>排练，同时通过‘点单派单’平台，邀请了专业的舞蹈老师进行指导。</p><p style="text-indent:2em;">“现在咱们舞蹈队队员达到了二十余人，大家排练节目都很积极。我们的节目已经在区里很多文艺汇演上表演过了，这多亏了李燕的帮忙。”宋芳珍感激地说。</p><p style="text-indent:2em;">为社区群众提供娱<a href='http://www.mojbuy.com/nj9juftbc/'><em>宝马平特论坛_蜘蛛资讯网</em></a>乐平台的同时，在李燕的主持下，实践站还成立了七支志愿服务队伍，其中将原有的恒安社区“1+X”搭把手志愿服务队融入实践站，由原来的的50名志愿者发展到现在的300名，所有在职党员、共建单位、青少年、社区民警都成了服务队的一员，<a href='http://jqe5u8v.mhdkc.top/58gw9/'><em>2019全年跑狗玄机图_蜘蛛资讯网</em></a>针对社区失独老人、残疾人等弱势群体开展送温暖、送服务、送真情等系列活动。</p><p style="text-indent:2em;">83岁的沈秀芬是恒安社区的老党员，多年独身生活的她一直是李燕重点关照的对象。自从李燕到恒安社区工作起，她就成了沈秀芬家的常客。逢年过节、下雨下雪她都要带人过来看望沈秀芬。</p><p style="text-indent:2em;">“沈阿姨这些年来挺不容易的，即使没有这份儿责任，我也应该帮扶她一下！有些事儿对我们年轻人来说很容易，但是对他们上了年纪的人来说却很难。让他们能够过一个安心的晚年，这挺好的。”李燕说道。</p><p style="text-indent:2em;">因为临近“七一”，采访当天，李燕带着社区工作者又一次走进了沈秀芬家。老太太的屋子收拾的干净利落，见到李燕她们进来，她立马拿出水果招呼大家吃。李燕和沈秀芬聊了聊近来的身体状况，叮嘱她要注意身体，随后组织大家抢着帮沈秀芬又仔细打扫了一遍屋子。</p><p style="text-indent:2em;">“小李人不错，对我们这群老人家很好。现在咱们社区新时代文明实践站几支志愿服务队都帮助过我，我就是年纪大了，如果年轻一些，会跟他们一起做活动呢。”沈秀芬笑着说。</p><p style="text-indent:2em;">从沈秀芬家出来已经临近中午，而下午还有外省的<a href='http://7tt9o9n3.mhmi.top/o2s5/'><em>2019马经挂牌系列e_蜘蛛资讯网</em></a>参观团来基地学习，李燕匆匆巴拉了几口饭，便又忙活起讲解的事情。对于她来说这是工作的常态，忙起来常常顾不上家里。由于爱人在市区工作，孩子只能由她照顾，对于孩子，她满是亏欠。不过让她欣慰的是，儿子很懂事，对李燕的工作很支持。</p><p style="text-indent:2em;">“冬天的时候，有一次早上我室外执勤，当时天色还有些黑，我儿子非要跟我一起去。他说‘妈妈是大志愿者，我是小志愿者。’在学校，他也总是做一些帮助同学的事儿，可能潜移默化间，他受我的一些影响吧。”说起儿子李燕眼神里流露出慈爱。</p><p style="text-indent:2em;">现在，日常工作之余，李燕还做了一个计划：把恒安社区的新时代文明实践站的内容再丰富健全一些，让社区老百姓(603883)能获得更多的便利，让区内外来基地的参观团队能收获更多的经验。为此，她一直都在默默努力着。（通讯员 张宏民）</p></p><p>&nbsp;&nbsp;&nbsp;&nbsp;</p><p>&nbsp;&nbsp;&nbsp;&nbsp;			</p><p>&nbsp;&nbsp;&nbsp;&nbsp;</p><p>&nbsp;&nbsp;&nbsp;&nbsp;			<div style="text-align:right;font-size:12px">（责任编辑： HN666）</div></p></p></p></div>
		</div>
    </div>
  </div>
</div>

<div class="foot">
	<div id="links"><a href='http://www.www-4227c.com' target='_blank'>宝马论坛管家婆四肖期期准中特网</a>&nbsp;&nbsp;<a href='http://www.bbkgolf.com' target='_blank'>www85777王中王资料</a>&nbsp;&nbsp;<a href='http://www.www200605.com' target='_blank'>www.950999.com</a>&nbsp;&nbsp;<a href='http://www.www-4394b.com' target='_blank'>静心阁高手论坛</a>&nbsp;&nbsp;<a href='http://www.www-281114.com' target='_blank'>香港马会资料大全</a>&nbsp;&nbsp;<a href='http://www.kj717.net' target='_blank'>www.kj717.com</a>&nbsp;&nbsp;<a href='http://www.www-614411.com' target='_blank'>王中王开奖记录</a>&nbsp;&nbsp;<a href='http://www.pt606.com' target='_blank'>香港白小姐一码中特料</a>&nbsp;&nbsp;<a href='http://www.cjgdj.com' target='_blank'>www.130355.com</a>&nbsp;&nbsp;<a href='http://www.www56012a.com' target='_blank'>港最快开奖现场2019结果</a>&nbsp;&nbsp;</div>
	<div class="foot_c">
		Copyright ＠ 2016-2018 蜘蛛资讯网 版权所有
	</div>
</div>



<script src='/static/bd_tui.js'></script>
</body>
</html>
<pre class='wup2c'></pre><pre class='dxcq8'><span class='a7w'><i class='5acnjnyge'></i></span></pre><b class='erujyc'></b><em class='cgkq'></em><pre class='fhvjxzmc'></pre><b class='fxdc25f'></b><span class='fkz1wk0qh'><b class='zy32gf8l'><i class='xt6z'></i></b></span><em class='daw3'></em><em class='kxi3'></em><em class='p0zbd'></em><pre class='5ft5'></pre><b class='ft9boik'></b><span  class='oca90'></span><span class='71fh5f'><em class='owkfu'></em></span><pre class='2e3v82hkj'></pre><b class='9jwk'></b><span class='lgewu'><em class='99vb'></em></span><em class='vjgplc'></em><em class='2o8'></em><pre class='o3kkd9j'><span class='99a2t'><i class='usxvfc'></i></span></pre><pre class='iizo'></pre><i class='fii8a'></i><b class='zg9'></b>
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