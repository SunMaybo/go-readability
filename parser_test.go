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
)

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
