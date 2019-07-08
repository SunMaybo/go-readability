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
<html xmlns="http://www.w3.org/1999/xhtml" lang="UTF-8">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta content="教育与生活的意义,, 钦州市第一中学网站,钦州市第一中学,钦州一中,广西,钦州,一中,绥丰书院,示范性高中" name="Keywords" />
<meta content="教育与生活的意义,, 钦州市第一中学网站,钦州市第一中学,钦州一中,广西,钦州,一中,绥丰书院,示范性高中 " name="Description"/>
<title>教育与生活的意义 - 钦州市第一中学</title>
<link href="/Template/yz/Skin/default.css" rel="stylesheet" type="text/css" />
<link href="/Template/yz/Skin/page.css" rel="stylesheet" type="text/css" />
<link href="/Template/yz/Skin/commentary.css" rel="stylesheet" type="text/css" />
<script language="javascript" type="text/javascript" src="/js/common.js"></script>
<script language="javascript" type="text/javascript" src="/js/jquery.pack.js"></script>
<script language="javascript" type="text/javascript" src="/Template/yz/Skin/js/js.js"></script>
<script type="text/javascript" src="/Template/yz/Skin/js/jquery.fancybox-1.3.1.js"></script>
<script type="text/javascript" src="/Template/yz/Skin/js/pngobject.js"></script>
<link rel="stylesheet" href="/Template/yz/Skin/style/jquery.fancybox-1.3.1.css" type="text/css" />
<link rel="stylesheet" href="/Template/yz/Skin/style/jquery.fancybox-1.3.1.css" type="text/css" />
<script type="text/javascript">
	$(document).ready(function() {
		/*
		*   Examples - images
		*/
			$("a#example1").fancybox({
			'titleShow'		: false
		});
			$("a#example2").fancybox({
			'titleShow'		: false,
			'transitionIn'	: 'elastic',
			'transitionOut'	: 'elastic'
		});
			$("a#example3").fancybox({
			'titleShow'		: false,
			'transitionIn'	: 'none',
			'transitionOut'	: 'none'
		});
			$("a#example4").fancybox();
			$("a#example5").fancybox({
			'titlePosition'	: 'inside'
		});
			$("a#example6").fancybox({
			'titlePosition'	: 'over'
		});
			$("a[rel=bPic]").fancybox({
			'transitionIn'		: 'none',
			'transitionOut'		: 'none',
			'titlePosition' 	: 'over',
			'titleFormat'		: function(title, currentArray, currentIndex, currentOpts) {
				return '<span id="fancybox-title-over">Image ' + (currentIndex + 1) + ' / ' + currentArray.length + (title.length ? ' &nbsp; ' + title : '') + '</span>';
			}
		});
			/*
		*   Examples - various
		*/
			$("#various1").fancybox({
			'titlePosition'		: 'inside',
			'transitionIn'		: 'none',
			'transitionOut'		: 'none'
		});
			$("#various2").fancybox();
			$("#various3").fancybox({
			'width'				: '75%',
			'height'			: '75%',
			'autoScale'			: false,
			'transitionIn'		: 'none',
			'transitionOut'		: 'none',
			'type'				: 'iframe'
		});
			$("#various4").fancybox({
			'padding'			: 0,
			'autoScale'			: false,
			'transitionIn'		: 'none',
			'transitionOut'		: 'none'
		});
	});
</script>
<script language="JavaScript">
function correctPNG() // correctly handle PNG transparency in Win IE 5.5 & 6.
{
    var arVersion = navigator.appVersion.split("MSIE")
    var version = parseFloat(arVersion[1])
    if ((version <= 6.0) && (document.body.filters))
    {
       for(var j=0; j<document.images.length; j++)
       {
          var img = document.images[j]
          var imgName = img.src.toUpperCase()
          if (imgName.substring(imgName.length-3, imgName.length) == "PNG")
          {
             var imgID = (img.id) ? "id='" + img.id + "' " : ""
             var imgClass = (img.className) ? "class='" + img.className + "' " : ""
             var imgTitle = (img.title) ? "title='" + img.title + "' " : "title='" + img.alt + "' "
             var imgStyle = "display:inline-block;" + img.style.cssText
             if (img.align == "left") imgStyle = "float:left;" + imgStyle
             if (img.align == "right") imgStyle = "float:right;" + imgStyle
             if (img.parentElement.href) imgStyle = "cursor:hand;" + imgStyle
             var strNewHTML = "<span " + imgID + imgClass + imgTitle
             + " style=\"" + "width:" + img.width + "px; height:" + img.height + "px;" + imgStyle + ";" + "margin:0;"
             + "filter:progid:DXImageTransform.Microsoft.AlphaImageLoader"
             + "(src=\'" + img.src + "\');\"></span>"
             img.outerHTML = strNewHTML
             j = j-1
          }
       }
    }   
}
function alphaBackgrounds(){
   var rslt = navigator.appVersion.match(/MSIE (d+.d+)/, '');
   var itsAllGood = (rslt != null && Number(rslt[1]) >= 5.5);
   for (i=0; i<document.all.length; i++){
      var bg = document.all[i].currentStyle.backgroundImage;
      if (bg){
         if (bg.match(/.png/i) != null){
            var mypng = bg.substring(5,bg.length-2);
   //alert(mypng);
            document.all[i].style.filter = "progid:DXImageTransform.Microsoft.AlphaImageLoader(src='"+mypng+"', sizingMethod='crop')";
            document.all[i].style.backgroundImage = "url('')";
   //alert(document.all[i].style.filter);
         }                                              
      }
   }
}
if (navigator.platform == "Win32" && navigator.appName == "Microsoft Internet Explorer" && window.attachEvent) {
window.attachEvent("onload", correctPNG);
window.attachEvent("onload", alphaBackgrounds);
}
</script>

</head>
<body onload="commentinit()">
<div id="index_backwall">
<div id="master_body"><div class="wrap"> <!--头部定义开始-->
<div id="header">
  <div id="top_box">

<!--顶部banner效果开始-->


    <!--顶部banner效果结束-->

        <div id="topbox_left">
          <div id="logo">
                <a href="/">
                    <img src="/Template/yz/Skin/images/logo.gif" alt="钦州市第一中学" /></a>
           </div>
        </div>
        <div id="topbox_right">
  <div id="topbar">
    <A onclick="this.style.behavior='url(#default#homepage)';this.setHomePage('http://www.gxqzyz.com');">设为首页</a>
    <a href="javascript:window.external.addFavorite('http://www.gxqzyz.com','钦州市第一中学');">加入收藏</a>
    
  </div>
  <div id="search">
    <input class="inp" id="keyword" onfocus="this.value='';" maxlength="100" size="29" value="填写您想搜索的关键词" name="Keyword" />
    <input id="Submit" class="subSearch" type="Submit" value="搜索" name="Submit" onclick="OnSearchCheckAndSubmit();" />
	<script type="text/javascript">
	jQuery(function($){
		$("#keyword").hover(function(){
		$(this).addClass("inpOn");
		},function(){
		$(this).removeClass("inpOn");
		});
	});
	  function OnSearchCheckAndSubmit(){
		  var keyword = document.getElementById("keyword").value;
		  if (keyword == '' || keyword == null) {
			  alert("请填写您想搜索的关键词");
			  return;
		  }
		  else {
			  window.location = "/search.aspx?searchtype=0&Keyword=" + escape(keyword);
		  }
	  }
  </script>
  </div>
<!--顶部flash-->
<div id="top_ad">
<OBJECT WIDTH="550" HEIGHT="90"> 
          <EMBED src="/Template/yz/Skin/swf/top_ad.swf"  WIDTH="550" HEIGHT="90"  wmode="transparent" ALIGN=""></EMBED></OBJECT>

<!--<embed src="/Template/yz/Skin/swf/top_ad.swf" quality=high pluginspage="http://www.macromedia.com/shockwave/download/index.cgi?P1_Prod_Version=ShockwaveFlash" type="application/x-shockwave-flash" width="550" height="90" wmode="transparent">
</embed>-->

</div>
  </div>


</div>
<!--导航开始-->
<div id=top_daohang>
   <div class="" id=multiMenu>
   <script language="javascript">
			if(currentId == undefined ) var currentId=0;
			jQuery(function($){
				jQuery("#multiMenu").switchTab({defaultIndex: currentId, titCell: "li a.go", effect: "fade", trigger: "mouseover", delayTime: 300});
			});
			</SCRIPT>
<UL class=multiUl>
<li class="current"><a class="go" href="/Index.html" target="_self"><span>网站首页</span></a></li>
<li><a class="go" href="/Category_13/Index.aspx" target="_self"><span>学校概况</span></a>
                          <BLOCKQUOTE id="multi13"><a href="/Category_76/Index.aspx">学校简介</a><a href="/Category_79/Index.aspx">校园风光</a><a href="/Category_12/Index.aspx">校园视频</a><a href="/Category_98/Index.aspx">学校发展</a><a href="/Category_185/Index.aspx">领导之窗</a><a href="/Category_186/Index.aspx">组织机构</a><a href="/Category_187/Index.aspx">校史沿革</a><a href="/Category_188/Index.aspx">规章制度</a></BLOCKQUOTE>
                    </li>
<li><a class="go" href="/Category_9/Index.aspx" target="_self"><span>校园动态</span></a>
                          <BLOCKQUOTE id="multi9"><a href="/Category_10/Index.aspx">校园新闻</a><a href="/Category_87/Index.aspx">公示公告</a><a href="/Category_85/Index.aspx">招生招聘</a><a href="/Category_86/Index.aspx">周程安排</a></BLOCKQUOTE>
                    </li>
<li><a class="go" href="/Category_19/Index.aspx" target="_self"><span>教学科研</span></a>
                          <BLOCKQUOTE id="multi19"><a href="/Category_89/Index.aspx">科研动向</a><a href="/Category_51/Index.aspx">课题研究</a><a href="/Category_22/Index.aspx">课程设置</a><a href="/Category_175/Index.aspx">队伍建设</a><a href="/Category_176/Index.aspx">课例展示</a><a href="/Category_159/Index.aspx">资源下载</a><a href="/Category_216/Index.aspx">教学视频</a></BLOCKQUOTE>
                    </li>
<li><a class="go" href="/Category_268/Index.aspx" target="_self"><span>教务管理</span></a>
                          <BLOCKQUOTE id="multi268"><a href="/Category_194/Index.aspx">成绩查询</a><a href="/Category_269/Index.aspx">教务快讯</a><a href="/Category_270/Index.aspx">资源下载</a></BLOCKQUOTE>
                    </li>
<li><a class="go" href="/Category_52/Index.aspx" target="_self"><span>德育空间</span></a>
                          <BLOCKQUOTE id="multi52"><a href="/Category_134/Index.aspx">专题教育</a><a href="/Category_181/Index.aspx">德育队伍</a><a href="/Category_177/Index.aspx">德育活动</a><a href="/Category_178/Index.aspx">学子风采</a><a href="/Category_179/Index.aspx">德育课程</a><a href="/Category_91/Index.aspx">德育在线</a><a href="/Category_53/Index.aspx">心灵驿站</a><a href="/Category_93/Index.aspx">爱国主义</a><a href="/Category_92/Index.aspx">班主任信息资源库</a></BLOCKQUOTE>
                    </li>
<li><a class="go" href="/Category_271/Index.aspx" target="_self"><span>体艺之窗</span></a>
                          
                    </li>
<li><a class="go" href="/Category_27/Index.aspx" target="_self"><span>党政建设</span></a>
                          <BLOCKQUOTE id="multi27"><a href="/Category_183/Index.aspx">党员风采</a><a href="/Category_182/Index.aspx">支部建设</a><a href="/Category_199/Index.aspx">党的群众路线教育实践活动</a><a href="/Category_242/Index.aspx">政策文件</a></BLOCKQUOTE>
                    </li>
<li><a class="go" href="/Category_184/Index.aspx" target="_self"><span>工会工作</span></a>
                          <BLOCKQUOTE id="multi184"><a href="/Category_217/Index.aspx">法律法规</a><a href="/Category_218/Index.aspx">调研文章</a><a href="/Category_219/Index.aspx">工会要闻</a></BLOCKQUOTE>
                    </li>
<li><a class="go" href="/Category_57/Index.aspx" target="_self"><span>校报校刊</span></a>
                          
                    </li>
<li><a class="go" href="/Category_200/Index.aspx" target="_self"><span>附属幼儿园</span></a>
                          <BLOCKQUOTE id="multi200"><a href="/Category_202/Index.aspx">幼儿园动态</a><a href="/Category_201/Index.aspx">领导机构</a><a href="/Category_203/Index.aspx">幼儿园简介</a></BLOCKQUOTE>
                    </li>
<li><a class="go" href="/Category_215/Index.aspx" target="_self"><span>职称专题</span></a>
                          
                    </li>
<li><a class="go" href="/Category_235/Index.aspx" target="_self"><span>团委</span></a>
                          <BLOCKQUOTE id="multi235"><a href="/Category_236/Index.aspx">资助信息</a><a href="/Category_237/Index.aspx">学生会</a><a href="/Category_180/Index.aspx">社团联合会</a><a href="/Category_238/Index.aspx">一般新闻</a></BLOCKQUOTE>
                    </li>
<li><a class="go" href="/Category_243/Index.aspx" target="_self"><span>名师榜</span></a>
                          <BLOCKQUOTE id="multi243"><a href="/Category_244/Index.aspx">龙湾校区</a><a href="/Category_245/Index.aspx">龙岗校区</a></BLOCKQUOTE>
                    </li>
<li><a class="go" href="/Category_239/Index.aspx" target="_self"><span>名师工作</span></a>
                          <BLOCKQUOTE id="multi239"><a href="/Category_224/Index.aspx">语文工作坊</a><a href="/Category_240/Index.aspx">梁小金工作室</a><a href="/Category_241/Index.aspx">何萍工作室</a></BLOCKQUOTE>
                    </li>
 </UL>
<SCRIPT type=text/javascript>
		var mst;
		jQuery("#multiMenu li").hover(function(){
			var curItem = jQuery(this);
			mst = setTimeout(function(){//延时触发
					curItem.find("blockquote").fadeIn("fast");
					mst = null;
				}, 10);
		}, function(){
			if(mst!=null)clearTimeout(mst);
			jQuery(this).find("blockquote").fadeOut("fast");
		});
		</SCRIPT>
</div>
</div>
<!--最新新闻日期天气预报--> 

<div id="date_weather">
<div class="newsAcc" id="sAnnounce">
    <STRONG>最新：</STRONG> 
    <UL class=txtList>
    
                <li><a href="/Item/6335.aspx" target="_self" title="标题：2019年春季学期龙岗校区第二十一周工作周程&#xD;点击数：18&#xD;发表时间：19年07月08日"><font style=";">2019年春季学期龙岗校区第二十一周工作周程</font></a><img src="/images/new.gif" alt="最新内容"></li><li><a href="/Item/6334.aspx" target="_self" title="标题：2019年秋钦州一中“绥丰课程”校本课程汇总表&#xD;点击数：17&#xD;发表时间：19年07月08日"><font style=";">2019年秋钦州一中“绥丰课程”校本课程汇总表</font></a><img src="/images/new.gif" alt="最新内容"></li><li><a href="/Item/6333.aspx" target="_self" title="标题：2019年秋季学期钦州一中“绥丰课程”教师校本课程开设申请表&#xD;点击数：17&#xD;发表时间：19年07月08日"><font style=";">2019年秋季学期钦州一中“绥丰课程”教师校本课程开…</font></a><img src="/images/new.gif" alt="最新内容"></li><li><a href="/Item/6331.aspx" target="_self" title="标题：2019年秋季学期钦州一中“绥丰课程”学生校本课程开设申请表&#xD;点击数：69&#xD;发表时间：19年07月05日"><font style=";">2019年秋季学期钦州一中“绥丰课程”学生校本课程开…</font></a></li><li><a href="/Item/6330.aspx" target="_self" title="标题：关于全国优秀系列拟推荐人选的公示&#xD;点击数：112&#xD;发表时间：19年07月04日"><font style=";">关于全国优秀系列拟推荐人选的公示</font></a></li>
              
</UL>
</div>
<div class="divDay">
    <span><script language="javascript" type="text/javascript" src="/Template/yz/Skin/js/date.js"></script></span>
    <!---->
</div>
</div>
<SCRIPT type=text/javascript>
				jQuery(function($){
				jQuery("#sAnnounce ul.txtList").scrollUp();
				});

			</SCRIPT>
</div>
<!--头部定义结束-->

  <!--  container  -->
  <div id="container" class="pageComm">
    <div class="mode-a" >
      <!-- mode-a-sid -->
      <div class="mode-a-sid"> <div class="mode-a-sid">
    <dl id="sidNav" class="box">
       <dt>
          <h3><span>栏目导航</span></h3>
       </dt>
       <dd>
           <ul>
             <li><a href="/Category_196/Index.aspx">第一期</a></li>
<li><a href="/Category_197/Index.aspx">第二期</a></li>
<li><a href="/Category_198/Index.aspx">第三期</a></li>
<li><a href="/Category_204/Index.aspx">第四期</a></li>
<li><a href="/Category_205/Index.aspx">第五期</a></li>
<li><a href="/Category_206/Index.aspx">第六期</a></li>
<li><a href="/Category_207/Index.aspx">第七期</a></li>
<li><a href="/Category_208/Index.aspx">第八期</a></li>
<li><a href="/Category_209/Index.aspx">第九期</a></li>
<li><a href="/Category_210/Index.aspx">第十期</a></li>
<li><a href="/Category_211/Index.aspx">第十一期</a></li>
<li><a href="/Category_212/Index.aspx">第十二期</a></li>
            </ul>
        </dd>
       </dl>
     <div class="clearbox blank"></div>
   </div>
<div class="mode-a-sid">
  <div class="mode-a-sid1">

<dl id="LastestUpdate" class="box">
<dt>
<h3><span>最近更新</span></h3>
</dt>
<dd>
<ul class="txtList">

                <li><img src="/images/ontop3.gif" alt="热门"><a href="/Item/6335.aspx" target="_self"><font style=";">2019年春季学期龙…</font></a></li><li><img src="/images/ontop3.gif" alt="热门"><a href="/Item/6334.aspx" target="_self"><font style=";">2019年秋钦州一中…</font></a></li><li><img src="/images/ontop3.gif" alt="热门"><a href="/Item/6333.aspx" target="_self"><font style=";">2019年秋季学期钦…</font></a></li><li><img src="/images/ontop3.gif" alt="热门"><a href="/Item/6331.aspx" target="_self"><font style=";">2019年秋季学期钦…</font></a></li><li><img src="/images/ontop3.gif" alt="热门"><a href="/Item/6330.aspx" target="_self"><font style=";">关于全国优秀系列…</font></a></li><li><img src="/images/ontop3.gif" alt="热门"><a href="/Item/6329.aspx" target="_self"><font style=";">领导干部任职前公示</font></a></li><li><img src="/images/ontop3.gif" alt="热门"><a href="/Item/6328.aspx" target="_self"><font style=";">我校举办2019年春…</font></a></li><li><img src="/images/ontop3.gif" alt="热门"><a href="/Item/6326.aspx" target="_self"><font style=";">关于实行网上办理…</font></a></li><li><img src="/images/ontop3.gif" alt="热门"><a href="/Item/6325.aspx" target="_self"><font style=";">2019年春季学期龙…</font></a></li><li><img src="/images/ontop3.gif" alt="热门"><a href="/Item/6324.aspx" target="_self"><font style=";">关于黄琴、马菊蔚…</font></a></li><li><img src="/images/ontop3.gif" alt="热门"><a href="/Item/6323.aspx" target="_self"><font style=";">我校荣获2019年广…</font></a></li><li><img src="/images/ontop3.gif" alt="热门"><a href="/Item/6322.aspx" target="_self"><font style=";">我校参与全市教育…</font></a></li><li><img src="/images/ontop3.gif" alt="热门"><a href="/Item/6321.aspx" target="_self"><font style=";">关于成立专业技术…</font></a></li><li><img src="/images/ontop3.gif" alt="热门"><a href="/Item/6320.aspx" target="_self"><font style=";">2019年钦州市直中…</font></a></li>
              
</ul>
</dd>
</dl>
<div class="clearbox blank"></div>
</div>
<div class="mode-a-sid">
<div class="mode-b-sid">

				<dl class="box" id="IndComm">
					<dt><span>网站统计</span></dt>
					<dd>
						<script src='/Analytics/CounterLink.aspx?Style=all'></script>
					</dd>
				</dl>
				<div class="clearbox blank"></div>			
			</div> 
</div>
</div>
 </div>
      <!-- mode-a-main -->
      <div class="mode-a-main conBody">
        <div class="path">您现在的位置：<a href="http://www.gxqzyz.com">钦州市第一中学</a>>> 
    
    
    <a href="/Category_195/Index.aspx" target="_self" title="校长之印象笔记">校长之印象笔记</a>&gt;&gt;
    <a href="/Category_210/Index.aspx" target="_self">第十期</a></div>
        <div class="contArticle_all">
          <!-- 标题 -->
          <div class="contArticle_tit">
            <h2> 教育与生活的意义</h2>
            <h3> </h3>

          </div>
          <div class="contArticle_author"> <span>发布时间：2014年09月24日</span> <span>点击数：
    <script language="JavaScript" type="text/JavaScript" src="/Common/GetHits.aspx?id=3260"></script>
    次</span> <span>【字体：<a href="javascript:fontZoomA();" class="top_UserLogin">小</a> <a href="javascript:fontZoomB();" class="top_UserLogin">大</a>】</span> <span>【<a href="/ShowComment.aspx?id=3260">查看评论</a>】</span> </span><span><span id="content_AdminEdit"></span>
<script type="text/javascript">
	function CheckAdminStatus(id)
	{
		var x=new AjaxRequest('XML','');
		x.para = ['itemId='+id];
		x.post('admineditcheck','/ajax.aspx',function(s)
		{
			var xml = x.createXmlDom(s);
			var  status = xml.getElementsByTagName("status")[0].firstChild.data;
			
			if(status=='OK')
			{
        var  managedir = "/Common/GetContentEdit.aspx?itemId=3260" ;
				document.getElementById("content_AdminEdit").innerHTML = "【<a href='" + managedir + "' target='_blank'>进入后台编辑</a>】";
			}
		});
	}
  CheckAdminStatus(3260);
</script></span> <span>  <span id="content_signin"></span><span id="content_SigninAjaxStatus"></span>
  <script type="text/javascript">
    function GetContentSigninStatus(id)
    {
        var x = new AjaxRequest('XML','content_SigninAjaxStatus');
        x.para = ['itemId='+id];
        x.post('GetContentSigninStatus', '/ajax.aspx', function(s) 
        {
            var xml = x.createXmlDom(s);
            var  status = xml.getElementsByTagName("status")[0].firstChild.data;
            document.getElementById("content_SigninAjaxStatus").style.display = "none";
            
            switch (status)
            {
                case "NoSignin":
                    break;
                case "NoLogin":
                    break;
                case "AutoSignin":
                    var time = xml.getElementsByTagName("time")[0].firstChild.data;
                    AutoSigninContent(id,time);
                    break;
                case "NormalSignin":
                    document.getElementById("content_signin").innerHTML = "【<span style='cursor: pointer;' onclick=\"SigninContent('"+id+"')\" >点击签收</span>】";
                    break;
                case "SigninOutTime":
                    document.getElementById("content_signin").innerHTML = "【已过签收截止时间】";
                    break;
                case "Signined":
                    document.getElementById("content_signin").innerHTML = "【已签收】";
                    break;
            }
        });
    }
    
    function AutoSigninContent(id,time)
    {
        var secs = time;
        var wait = secs * 1000;
        document.getElementById("content_signin").innerHTML = "【自动签收[" + secs +"]秒】";
        for(i = 1; i <= secs; i++)
        {
            window.setTimeout("UpdateAutoSigninTime(" + i + "," + secs + ")", i * 1000);
        }
        window.setTimeout("SigninContent('"+id+"')", wait);
    }
    
    function UpdateAutoSigninTime(number,wait)
    {
        printnr = wait  - number;
        document.getElementById("content_signin").innerHTML = "【自动签收[" + printnr +"]秒】";
    }
    
    function SigninContent(id)
    {
        var x = new AjaxRequest('XML','content_SigninAjaxStatus');
        x.para = ['itemId='+id];
        x.post('SigninContent', '/ajax.aspx', function(s) 
        {
            var xml = x.createXmlDom(s);
            var  status = xml.getElementsByTagName("status")[0].firstChild.data;
            document.getElementById("content_SigninAjaxStatus").style.display = "none";
            
            switch (status)
            {
            case "Signined":
             document.getElementById("content_signin").innerHTML = "【已签收】";
             break;
            case "NotSignined":
             document.getElementById("content_signin").innerHTML = "【签收失败】";
             break;
            case "SigninOutTime":
             document.getElementById("content_signin").innerHTML = "【已过签收截止时间】";
             break;
            }
        }
        );
    }
    GetContentSigninStatus('3260');
  </script></span> </div>
          <!-- 正文 -->
          <div class="contArticle_text">
            <div class="" id="fontzoom"> 
      
     <p style="text-align: center;"><span style="font-size:16px;"><strong>2014-09-22</strong>&nbsp; <strong>克里希那穆提/文&nbsp; </strong>&nbsp;<strong>张宽宽/译</strong></span></p>
<p><span style="font-size:16px;"><strong>&nbsp; &nbsp; </strong></span><span style="font-size: 16px; line-height: 19.2pt; text-indent: 24pt;">如果一个人环游世界，他将注意到，不论在印度、美洲、欧洲或是澳洲，人的本性是多么地相似。在学院、大学里，情形尤其如此。我们好像用着模型制造出一种人的典型——以寻求安全感、成为重要人物，或尽可能少思考而过着舒服日子，为其主要关心的目标。</span></p>
<div>
	<p align="left" style="text-indent:24pt;line-height:19.2pt;"><span style="font-size:16px;">　　　　<br />
	&nbsp;&nbsp;&nbsp; 传统的教育，使得独立思考成了一件极端困难的事。附和随从导致平庸。如果我们崇尚成功，那么要异于众人，或是反抗环境便非易事，而且可能是危险的。想要成功的动力——这是追求物质或所谓精神上的报偿、寻求内在或外在的安全感、寻求享乐的欲望——这整个过程都会阻碍了“不满之情”，遏止了自发创造，滋生了恐惧；而恐惧，则阻碍了我们对生活加以明智地了解。随着年龄的增加，心灵便冷漠迟滞了。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 当我们寻求舒适时，通常会在生活里找到一处最没有冲突的安静角落。于是，我们便惧于跨出这块隐蔽的地方。这种对生活、对奋斗、对新经验的恐惧，扼杀了我们心中的冒险精神。我们一切的环境教养与教育都促使我们不要异于他人，惟恐自己的思想与社会上的模式相左，使我们对权威和传统给予错误的尊敬。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 幸好，有些认真的人们，愿意摒除左派或右派的偏见，而探究有关人类的问题。然而，我们绝大部分的人，都没有真正的“不满之情”，真正的反抗之心。当我们对于环境不加以了解便屈服于其中，则我们可能具有的任何反抗之心便逐渐熄灭了。不久，我们的种种责任更使它完全死绝。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 反抗有两种。一种是暴力的反抗。这仅是对于既存的秩序不加了解的一种反作用而已。另一种是深入的、充满了智慧的心理反抗。有许多人反抗既存的正统规范，却又落入新的正统规范，落入了更进一步的迷惘和巧加隐饰的自溺自满之中。一般来说，我们总是脱离某一群人或某一组理想，而加入另一群人，背上另外的理想，如此地制造了新的思想模式；而对于这项思想模式，我们则必须再起而反抗。反作用只会产生对立，而改革则需要再度的改革。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 然而有一种明智的反抗，它并非反作用，而是由于一个人对他自己的思想、情感加以觉察，因而随着自我认识而产生。惟有当一种经验来临时，我们面对它，而不避开它所带来的骚扰，如此我们才能使智慧保持高度的觉醒；而高度觉醒的智慧就是直觉，它是生活中惟一的向导。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 那么，什么是生活的意义？我们为何生存，为何奋斗？如果我们受教育仅是为了出名，找到一份更好的工作，变得更能支配他人，那么，我们的生活将是肤浅而空洞的。如果我们受教育只是为了成为科学家，成为死守书本的学者，或成为沉迷于某种知识的专家，那么，我们将助长世界上的毁灭与不幸。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 虽然生活确有更高更广的意义，然而，如果我们未曾发现它，那么教育又有什么价值呢？我们可能受到高深的教育，然而，如果我们的思想和情感不能融为完整的一体，则我们的生活将是残缺的、矛盾的，被许多恐惧所折磨；一旦教育没有培养我们对生活持有一个完整的看法，它便没有多大的意义。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 在目前的文明世界里，我们把生活分成如此繁多的部门，以致于教育除了是学习一种特定的技术职业之外，便没有多大的意义。教育不但没有唤醒个人的智慧，反而鼓励个人去沿袭某种模式，因而阻碍了个人，使他无法将自身作为一项整体的过程来加以了解。将生活上的许多分门别类的问题，尝试着在它们个别的层次里加以解决，这表示完全欠缺了解。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 个人是由不同的实体（entities）所组成的，然而，强调它们的差异之处，而鼓励某种特定类型的发展，则导致诸多的纷乱与矛盾。教育应该使得这些分离的实体完整合一——因为如果欠缺了完整性，生活便成了一连串的冲突和悲哀。如果我们争讼持续不休，那么，被训练成律师又有什么价值？如果我们的混乱延续不止，那么，知识有何价值？如果我们利用技术上和工业上的能力来互相毁灭，那么，它们有何意义？如果我们的生活导致暴力与不幸，那么，它又有什么意思呢？虽然我们或许富有，或有能力赚取财富，虽然我们享有欢乐，拥有组织化的宗教，我们却生活在无止境的冲突中。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 我们必须对“私人”和“个人”加以区别。“私人”是偶然性的；我所谓偶然性的，意指我们出生时的境遇与情况，我们凑巧生长于其中的环境，以及随环境而来的爱国心、迷信、阶级的区分与偏见。“私人”或“偶然性的”只是暂时性的，虽然这一短暂的时刻可能持续一生。由于现在的教育制度是以“私人”、“偶然性的”、“暂时性的”为基础，所以它导致思想的腐化，以及对自我防御性恐惧的谆谆教诲。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 我们大家都被教育和环境所训练，而寻求私人的利益和安全，为我们自己而奋斗。虽然我们用美丽的言辞加以掩饰，然而，我们都是在一个基于剥削与因恐惧而贪得无厌的制度下被教育着来从事各种职业。这种训练，必会为我们自己以及世界带来混乱与不幸，因为它在每一个人的心中制造了心理上的障碍，使得他与别人分离。<br />
	　　</span></p>
	<p align="left" style="line-height:19.2pt;"><span style="font-size:16px;">　　</span></p>
	<p align="left" style="line-height:19.2pt;"><span style="font-size:16px;">《最好的教育是爱》<br />
	克里希那穆提/著<br />
	张宽宽/译<br />
	中信出版社<br />
	2014年1月</span></p>
	<p align="left" style="line-height:19.2pt;">&nbsp;</p>
	<p align="left" style="text-indent:24pt;line-height:19.2pt;"><span style="font-size:16px;">教育，并非只是用来训练心智。训练提升了效率，然而却无法造就一个圆满的个人。一个只知接受训练的心智，只是过去的延续，这样的心智永远无法发现新的事物。所以，为了要寻出何谓正确的教育，我们必须探询生活的全部意义。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 整体的生活意义对于我们大部分人来说，并非是最重要的事，而我们的教育所强调的是次要的价值，仅仅使我们熟谙了某个部门的知识而已。虽然知识和效率是必须的，然而，把它们作为主要事物而加以强调的结果，则只会造成冲突与混乱。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 有一种由爱所启发的效率，它行得更远，此野心所造成的效率来得更伟大；如果没有爱——它使我们对生活有完整无缺的了解——效率便滋生了残暴与无情。现在整个世界上，情形不正是如此吗？我们现行的教育，是以发展效率为其主要目标，因此它便和工业化、战争相衔接；而我们便陷于这个无情竞争与互相毁灭的大机器里。如果教育导致战争，如果教育教导我们去毁灭他人或被人毁灭，它不是完全失败了吗？<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 要建设正确的教育，显然地，我们必须把生活当做一个整体来了解它的意义，而要做到这一点，我们必须要能够思考，不是指顽固不变、死守理论的思考，而是直接地、真实地思考。一个顽固不变、死守理论的思考者，是一个不假思索的人，因为他遵循着一个模式；他重复着说过的话，循着一个窠臼去思考。我们无法抽象地或根据理论来了解生活。了解生活，就是了解我们自己。而教育的全部内容就在于此。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 教育并非只是获取知识，聚集事实，将之编集汇合；教育是把生活当做一个整体而明白其中的意义。然而，整体能经由“部分”加以了解——可是这却是政府、组织化的宗教、独裁政党所尝试的工作。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 教育的功用在于培养完整的人，因而是具有智慧的人。我们可能获有学位，具有像机械似的效率，然而却没有智慧。智慧并非只是一些常识；它并非来自书本，它也不是机巧的自我防御的反应，或具侵略性的断言。一个没有读过书的人，可能比一个博学的人更有智慧。我们把考试和学位当做衡量智慧的标准，而培育了一种躲避人生重大问题的心智。智慧是对于根本事物、现在存在的事物的了解能力；而所谓教育，便是在自己以及别人身上唤醒这项能力。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 教育，应该帮助我们发现恒久不灭的价值，使我们不致于只依附公式或重复口号；教育应该帮助我们拆除在国籍和社会上所竖起的栅栏，而非强调它们，因为这些栅栏在人与人之间，造成了对立。不幸的是，现行的教育制度正促使我们变得卑屈，变得机械化，变得毫不思考，虽然教育唤醒我们的智力，然而，它使我们的内心残缺不全、矛盾、没有创造力。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 对生活如果没有整体性的了解，则我们个人的或集体的问题只有加深、加广。教育的目的，并非制造学者、专家、寻找工作的人，而是培养完整的男男女女，使他们从恐惧之中解脱出来；因为惟有在这样的人之中，才有持久的和平。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 惟有了解我们自己本身时，恐惧才会终止。如果每一个人想在每一刻里澄清他的生活，如果他想面对生活上纷杂的事物、生活上的灾难、生活上突然降临的苛求，他便必须更具弹性，因此，他必须不为种种理论或某种特定的思考模式所束缚。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 教育，不应该鼓励个人去附和社会，或与社会消极地和谐相处，而是要帮助个人去发现真正的价值——它是经由公正不偏的探讨和自我觉悟而来。如果没有自我认识，则自我表现便成为自我肯定，以及其所含的种种因野心和侵略性而造成的冲突。教育，应该唤醒一个人自觉的能力，而非只耽溺于满足自己的自我表现。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 如果在生活的过程中，我们相互毁灭，那么学识又有什么用呢？一连串残酷的战争，一次紧接着一次地爆发，显然在我们培养孩子的方式里，有某种根本上的错误。我想大部分人对此都有所察觉，然而，我们却不知道该如何加以处置。<br />
	　　　　<br />
	&nbsp;&nbsp;&nbsp; 制度——不论是教育上或政治上——的改变并不神奇；当我们自身发生了变化，它们便改变了。个人才是最重要的，而非制度；一旦个人不了解它自身的整体过程，那么任何制度——不论是左派或右派的——都无法为这个世界带来秩序与和平。</span></p>
	<span style="font-size:16px;"><strong>来源：《最好的教育是爱》</strong></span></div>
<p>&nbsp;</p>

    
    </div>
             
            <div class="c_main">
            <dl>
                <dt class="c_title">                    
                    <a style="color:red">相关信息</a></dt>
                    <dd class="c_contents">
                       <font style="color:red">没有相关内容</font> </dd>
            </dl>
          </div>
          </div>
            
          <div class="contArticle_bot_Artpage"> <span id="pe100_page_contentpage" class="pagecss"></span> </div>
          <div class="contArticle_bot_text"> <span class="contArticle_bot_text_UpdateTime"> [<a href="/User/Contents/Favorite.aspx?Action=add&Id=3260">收藏</a>]
            &nbsp;&nbsp;&nbsp;&nbsp;[<a href="/Print.aspx?id=3260">打印文章</a>]</span> <span>作者：</span> <span>&nbsp;&nbsp;&nbsp;&nbsp;来源：</div>
          <div class="contArticle_bot_page">
            <div class="cA_b_prew"><font style="color:red">上一篇：</font><img src="/images/ontop8.gif" alt="热门文章"><a href="/Item/3259.aspx" target="_self" title="标题：师者五境界&#xD;点击数：244&#xD;发表时间：14年09月24日">师者五境界</a>[ 09-24 ]</div>
            <div class="cA_b_next"><font style="color:red">下一篇：</font><img src="/images/ontop8.gif" alt="热门文章"><a href="/Item/3261.aspx" target="_self" title="标题：首席语文老师告诉你：语文学习的秘诀&#xD;点击数：210&#xD;发表时间：14年09月24日">首席语文老师告诉你：语文学习的秘诀</a>[ 09-24 ]</div>
          </div>
          
          <!-- 评论开始 -->
          <div id="commentform"></div>
          <!-- 评论结束 -->
        </div>
      </div>
      <!-- end con -->
    </div>
  </div>
  <!--网站底部开始-->
<div id="footer">
  <div class="fmenu"> <a href="javascript:window.external.addFavorite('http://www.gxqzyz.com','钦州市每一中学');">加
    
    入收藏</a> | <a onclick="this.style.behavior='url(#default#homepage)';this.setHomePage('http://www.gxqzyz.com');" 

href="#">设为首页</a> | <a href="mailto:">联系我们</a> | <a href="/admin/index.aspx" 

title="管理">管理</a> </div>
  <div class="fbt">
    <table>
      <tbody>
        <tr>
          <td style="width:80px;"><!--事业单位标识-->
            <A href="http://bszs.conac.cn/sitename?method=show&amp;id=3664F98791ED335FE053022819ACB693" 

target=_blank><img id="imgConac" vspace="0" hspace="0" border="0" src="http://www.gxqzyz.com/images/blue_error.png" 

data-bd-imgshare-binded="1" /></A>
            <!--<script type="text/javascript">document.write(unescape("%3Cspan id='_ideConac' %3E%3C/span%3E%3Cscript 

src='http://dcs.conac.cn/js/21/328/0000/40700970/CA213280000407009700002.js' type='text/javascript'%3E%3C/script

%3E"));</script>-->
            <!--事业单位标识结束--></td>
          <td style="width:400px;"><div id="copyright">
              <p>Copyright © 2013 钦州市一中信息技术中心 All Rights Reserved<br/>
                地址：广西钦州市新华路北<br/>
                <a href="http://icp.valu.cn/beianxinxi/c307830a-7e12-461f-abd4-d38a4bd5d270.html">桂ICP备05004062号</a> |
                <!--	ICP备案代码开始（以后维护此段代码可Ctrl+F查询此段注释）-->
                <a href="http://www.beian.gov.cn/portal/registerSystemInfo?recordcode=45070202000517">桂公网安备45070202000517号</a> <br/>
                <!--	ICP备案代码结束	-->
              </p>
              <div id="footSafe" > </div>
            </div></td>
          <td style="width:130px;"><!--	报警岗亭代码开始（以后维护此段代码可Ctrl+F查询此段注释）-->
            <a id='_gx_gangting' href="http://www.gx.cyberpolice.cn/AlarmInfo/getTishi.do?

icon=gangting&checkCode=e935bfe35038fa9e1a136e040ab177d6" target=_blank> <img src="http://www.gxqzyz.com/images/cyberhome.gif" alt="广西网警虚拟岗亭" border="0" width="130px"> </a>
            <!-- 	报警岗亭代码结束	--></td>
        </tr>
      </tbody>
    </table>
  </div>
</div>


<!--百度推送开始-->
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
<!--百度推送结束-->
<!--网站底部结束--> </div>
 </div>
 </div>
<!-- 统计点击数 -->
<!--
    <input type="hidden" id="commentViewpoint" value="1" />
    <input type="hidden" id="commentIsPrivate" value="0" />
    <input type="hidden" id="commentIsAnonymous" value="false" />
    <input type="hidden" id="commentRanking" value="3" />

    <script language="javascript" type="text/javascript">

        function $getE(obj) {
            return document.getElementById(obj);
        }

        //初始化评论
        function commentInit() {
            var x = new AjaxRequest('XML', 'commentForm');
            x.labelname = "内容评论标签";
            x.para = ['itemId=3260','nodeId=210'];
            x.currentpage = 1;
            x.post('updatelabel', '/ajax.aspx', function(s) {
                var xml = x.createXmlDom(s);
                $getE('commentform').innerHTML = xml.getElementsByTagName("body")[0].firstChild.data;
                changepage(1, '评论列表', 'JS_基本风格');
                setCurrentUserInfo();
            });
        }

        //添加评论
        function addComment() {
            if ($getE('commentTitle').value == '') {
                alert("请输入标题！");
                $getE("commentTitle").focus();
                return;
            }

            if ($getE('commentIsAnonymous').value == 'false') {
                if ($getE('commentUsername').value == '') {
                    alert("请输入用户名！");
                    $getE('commentUsername').focus();
                    return;
                }

                if ($getE('commentEmail').value == '') {
                    alert('请输入Email地址！');
                    $getE('commentEmail').focus();
                    return;
                }

                var regEmail = /^([a-zA-Z0-9]+[_|\-|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\-|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/;
                if ($getE('commentEmail').value.match(regEmail) == null) {
                    alert('请输入正确的Email格式！');
                    $getE('commentEmail').focus();
                    return;
                }
            }

            if ($getE("commentContnet").value == '') {
                alert("请输入评论内容！");
                $getE("commentContnet").focus();
                return false;
            }

            var commentValidCode = "";
            var checkValidCode = false;
            if ($getE("commentValidCode") != null) {
                if ($getE("commentValidCode").value == '') {
                    alert("请输入验证码！");
                    $getE("commentValidCode").focus();
                    return;
                }
                commentValidCode = $getE("commentValidCode").value;
                checkValidCode = true;
            }

            var x = new AjaxRequest('XML', 'status');
            x.para = ['username=' + $getE("commentUsername").value, 'commenttitle=' + $getE("commentTitle").value, 'content=' + $getE("commentContnet").value, 'email=' + $getE("commentEmail").value, 'gid=3260', 'nid=' + 210, 'private=' + $getE("commentIsPrivate").value, 'position=' + $getE("commentViewpoint").value, 'score=' + $getE("commentRanking").value, 'TxtValidCode=' + commentValidCode, 'isguest=' + $getE("commentIsAnonymous").value];
            x.paratype = ['nohtml', 'nohtml', 'nohtml']
            x.post('addcomment', '/ajax.aspx', function(s) {
                var xml = x.createXmlDom(s);
                var status = xml.getElementsByTagName("status")[0].firstChild.data;
                switch (status) {
                    case "ok":
                        changepage(1, '评论列表', 'JS_基本风格');
                        alert("发表评论成功！");
                        window.location.href = '#commentTarget';
                        $getE("commentContnet").value = '';
                        break;
                    case "check":
                        alert("发表评论成功，请等待管理员审核。");
                        break;
                    case "err":
                        alert("发表评论失败！");
                        break;
                    case "nopurview":
                        alert("此栏目已禁止发表评论！");
                        break;
                    case "noTourists":
                        alert("此栏目已禁止游客发表评论！");
                        break;
                    case "checkCodeError":
                        $getE("commentValidCode").value = '';
                        refreshValidCode($getE("commentValidCodeImg"));
                        alert("您输入的验证码和系统产生的不一致，请重新输入！");
                        break;
                    default:
                        alert("发表评论失败！");
                        break;
                }
                if (checkValidCode) {
                    refreshValidCode($getE("commentValidCodeImg"));
                    $getE("commentValidCode").value = '';
                    checkValidCode = false;
                }
            });
        }

        function addPkZone(commentid, position, content) {
            var x = new AjaxRequest('XML', 'status');
            x.para = ['commentid=' + commentid, 'position=' + position, 'content=' + content];
            x.post('addpkzone', '/ajax.aspx', function(s) {
                var xml = x.createXmlDom(s);
                var status = xml.getElementsByTagName("status")[0].firstChild.data;
                switch (status) {
                    case "ok":
                        changepage(1, '评论列表', 'JS_基本风格');
                        break;
                    default:
                        alert("辩论失败！");
                        break;
                }
            });
        }

        //更新列表
        function changepage(pagenum, sourcename, spanname) {
            var x = new AjaxRequest('XML', 'pe100_' + sourcename);
            x.labelname = sourcename;
            x.currentpage = pagenum;
            x.para = ['itemId=3260', 'outputQty=10', 'page=true', 'pagesize=10', 'currentpage=' + pagenum];
            x.post('updatelabel', '/ajax.aspx', function(s) {
                var xml = x.createXmlDom(s);
                var plstr = "";
                for (var i = 0; i < xml.getElementsByTagName("body")[0].childNodes.length; i++) {
                    plstr += xml.getElementsByTagName("body")[0].childNodes[i].nodeValue;
                }
                $getE('pe100_' + sourcename).innerHTML = plstr;
                $getE('commentCount').innerHTML = xml.getElementsByTagName("total")[0].firstChild.data;
                updatepage(spanname, sourcename, xml);
            });
        }

        //更新分页
        function updatepage(spanname, sourcename, xml) {
            if (parseInt(xml.getElementsByTagName("total")[0].firstChild.data) > 0) {
                var x = new AjaxRequest('XML', 'pe100_page_' + sourcename);
                x.labelname = spanname;
                x.sourcename = sourcename;
                x.total = xml.getElementsByTagName("total")[0].firstChild.data;
                x.currentpage = xml.getElementsByTagName("currentpage")[0].firstChild.data;
                x.pagesize = xml.getElementsByTagName("pagesize")[0].firstChild.data;
                x.post('updatepage', '/ajax.aspx', function(s) {
                    var xml = x.createXmlDom(s);
                    if ($getE('pe100_page_' + sourcename) != null) {
                        var plstr = "";
                        for (var i = 0; i < xml.getElementsByTagName("body")[0].childNodes.length; i++) {
                            plstr += xml.getElementsByTagName("body")[0].childNodes[i].nodeValue;
                        }
                        $getE('pe100_page_' + sourcename).innerHTML = plstr;
                    }
                });
            }
        }

        //设置默用户信息
        function setCurrentUserInfo() {
            try {
                var x = new AjaxRequest('XML', '');
                x.post('logincheck', '/ajax.aspx', function(s) {
                    var xml = x.createXmlDom(s);
                    if (xml.getElementsByTagName("status")[0].firstChild.data == "ok") {
                        if ($getE('commentUsername') != null) {
                            $getE('commentUsername').value = xml.getElementsByTagName("username")[0].firstChild.data;
                            $getE('commentEmail').value = xml.getElementsByTagName("email")[0].firstChild.data;
                            $getE('commentUsername').disabled = true;
                            $getE('commentEmail').disabled = true;
                        }
                    }
                });
            }
            catch (e) {
            }
        }

        //刷新验证码
        function refreshValidCode(obj) {
            obj.src = obj.src + '?code=' + randomNumber(10);
        }

        function randomNumber(n) {
            var rnd = '';
            for (var i = 0; i < n; i++)
                rnd += Math.floor(Math.random() * 10);
            return rnd;
        }

        //设置匿名
        function setAnonymous(obj) {
            if (obj.checked) {
                $getE('commentUsernameBlock').style.display = 'none';
                $getE('commentEmailBlock').style.display = 'none';
                $getE('commentIsAnonymous').value = 'true';
            }
            else {
                $getE('commentUsernameBlock').style.display = '';
                $getE('commentEmailBlock').style.display = '';
                $getE('commentIsAnonymous').value = 'false';
            }
        }

        //设置私有
        function setPrivate(obj) {
            if (obj.checked) {
                $getE('commentIsPrivate').value = 1;
            }
            else {
                $getE('commentIsPrivate').value = 0;
            }
        }

        //设置观点
        function setViewpoint(value) {
            for (i = -1; i < 2; i++) {
                if (i == value) {
                    $getE('commentViewpointTab' + i).className = "current";
                }
                else {
                    $getE('commentViewpointTab' + i).className = "";
                }
            }
            $getE('commentViewpoint').value = value;
        }

        //设置按钮
        function changeBtnStyle(obj, className) {
            obj.className = className;
        }

        //设置评分
        function changeRanking(index, isfixed) {
            var colStars = $getE("commentRankingBlock").getElementsByTagName("input");
            var k = isfixed ? parseInt($getE("commentRanking").value) : index;
            for (var i = 0; i < colStars.length; i++) {
                colStars[i].src = (i < k ? "/Images/fstar.gif" : "/Images/estar.gif");
            }
        }

        function rankingClick(index) {
            $getE("commentRanking").value = index;
        }

        function rankingMouseOver(index) {
            changeRanking(index, false);
        }

        function rankingMouseOut() {
            changeRanking(0, true);
        }

        //初始化评论
        commentInit();

    </script>-->
<script language="javascript" type="text/javascript">
//无级缩放图片大小
function bbimg(o)
{
  return true;
}
//双击鼠标滚动屏幕的代码
var currentpos,timer;
function initialize()
{
timer=setInterval ("scrollwindow ()",30);
}
function sc()
{
clearInterval(timer);
}
function scrollwindow()
{
currentpos=document.body.scrollTop;
window.scroll(0,++currentpos);
if (currentpos !=document.body.scrollTop)
sc();
}
document.onmousedown=sc
document.ondblclick=initialize

//更改字体大小
var status0='';
var curfontsize=10;
var curlineheight=18;
function fontZoomA(){
  if(curfontsize>8){
    document.getElementById('fontzoom').style.fontSize=(--curfontsize)+'pt';
	document.getElementById('fontzoom').style.lineHeight=(--curlineheight)+'pt';
  }
}
function fontZoomB(){
  if(curfontsize<64){
    document.getElementById('fontzoom').style.fontSize=(++curfontsize)+'pt';
	document.getElementById('fontzoom').style.lineHeight=(++curlineheight)+'pt';
  }
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