package main

import (
	"fmt"
	"regexp"
)

var authorRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var publicRe = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+)<br>`)
var pagesRe =  regexp.MustCompile(`<span class="pl">页数:</span> 320<br>`)
var priceRe =  regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br>`)
var scoreRe =  regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
var introRe = regexp.MustCompile(`<div class="intro">[\d\D]*<p>([^<]+)</p></div>`)

func main()  {
	str := `<div id="content">
    
    <div class="grid-16-8 clearfix">
      
      <div class="article">



<div class="indent">
  <div class="subjectwrap clearfix">
    



<div class="subject clearfix">
<div id="mainpic" class="">

  

  <a class="nbg" href="https://img2.doubanio.com/view/subject/l/public/s11284102.jpg" title="霍乱时期的爱情">
      <img src="https://img2.doubanio.com/view/subject/s/public/s11284102.jpg" title="点击看大图" alt="霍乱时期的爱情" rel="v:photo" style="max-width: 135px;max-height: 200px;">
  </a>



</div>





<div id="info" class="">



    
    
  
    <span>
      <span class="pl"> 作者</span>:
        
            
            <a class="" href="/search/%E5%8A%A0%E8%A5%BF%E4%BA%9A%C2%B7%E9%A9%AC%E5%B0%94%E5%85%8B%E6%96%AF">[哥伦比亚] 加西亚·马尔克斯</a>
    </span><br>

    
    
  
    <span class="pl">出版社:</span> 南海出版公司<br>

    
    
  
    <span class="pl">出品方:</span>&nbsp;<a href="https://book.douban.com/series/39059?brand=1">新经典文化</a><br>

    
    
  

    
    
  
    <span class="pl">原作名:</span> El amor en los tiempos del cólera<br>

    
    
  
    <span>
      <span class="pl"> 译者</span>:
        
            
            <a class="" href="/search/%E6%9D%A8%E7%8E%B2">杨玲</a>
    </span><br>

    
    
  
    <span class="pl">出版年:</span> 2012-9-1<br>

    
    
  
    <span class="pl">页数:</span> 401<br>

    
    
  
    <span class="pl">定价:</span> 39.50元<br>

    
    
  
    <span class="pl">装帧:</span> 精装<br>

    
    
  
    <span class="pl">丛书:</span>&nbsp;<a href="https://book.douban.com/series/10489">新经典文库：加西亚·马尔克斯作品</a><br>

    
    
  
    
      
      <span class="pl">ISBN:</span> 9787544258975<br>


</div>

</div>
























    





<div id="interest_sectl" class="">
  <div class="rating_wrap clearbox" rel="v:rating">
    <div class="rating_logo">豆瓣评分</div>
    <div class="rating_self clearfix" typeof="v:Rating">
      <strong class="ll rating_num " property="v:average"> 9.0 </strong>
      <span property="v:best" content="10.0"></span>
      <div class="rating_right ">
          <div class="ll bigstar45"></div>
            <div class="rating_sum">
                <span class="">
                    <a href="comments" class="rating_people"><span property="v:votes">207481</span>人评价</a>
                </span>
            </div>


      </div>
    </div>
          
            
            
<span class="stars5 starstop" title="力荐">
    5星
</span>

            
<div class="power" style="width:64px"></div>

            <span class="rating_per">57.7%</span>
            <br>
            
            
<span class="stars4 starstop" title="推荐">
    4星
</span>

            
<div class="power" style="width:37px"></div>

            <span class="rating_per">33.5%</span>
            <br>
            
            
<span class="stars3 starstop" title="还行">
    3星
</span>

            
<div class="power" style="width:8px"></div>

            <span class="rating_per">7.9%</span>
            <br>
            
            
<span class="stars2 starstop" title="较差">
    2星
</span>

            
<div class="power" style="width:0px"></div>

            <span class="rating_per">0.6%</span>
            <br>
            
            
<span class="stars1 starstop" title="很差">
    1星
</span>

            
<div class="power" style="width:0px"></div>

            <span class="rating_per">0.2%</span>
            <br>
    </div>
</div>

  </div>
  





  
    
    <div id="interest_sect_level" class="clearfix">
        <a href="#" rel="nofollow" class="j a_show_login colbutt ll" name="pbtn-10594787-wish">
          <span>
            
<form method="POST" action="https://www.douban.com/register?reason=collectwish" class="miniform">
    <input type="submit" class="minisubmit j  " value="想读" title="">
</form>

          </span>
        </a>
        <a href="#" rel="nofollow" class="j a_show_login colbutt ll" name="pbtn-10594787-do">
          <span>
            
<form method="POST" action="https://www.douban.com/register?reason=collectdo" class="miniform">
    <input type="submit" class="minisubmit j  " value="在读" title="">
</form>

          </span>
        </a>
        <a href="#" rel="nofollow" class="j a_show_login colbutt ll" name="pbtn-10594787-collect">
          <span>
            
<form method="POST" action="https://www.douban.com/register?reason=collectcollect" class="miniform">
    <input type="submit" class="minisubmit j  " value="读过" title="">
</form>

          </span>
        </a>
      <div class="ll j a_stars">
          
    
    评价:
    <span id="rating"> <span id="stars" data-solid="https://img3.doubanio.com/f/shire/5a2327c04c0c231bced131ddf3f4467eb80c1c86/pics/rating_icons/star_onmouseover.png" data-hollow="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" data-solid-2x="https://img3.doubanio.com/f/shire/7258904022439076d57303c3b06ad195bf1dc41a/pics/rating_icons/star_onmouseover@2x.png" data-hollow-2x="https://img3.doubanio.com/f/shire/95cc2fa733221bb8edd28ad56a7145a5ad33383e/pics/rating_icons/star_hollow_hover@2x.png">

            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-10594787-1">
            <img src="https://img3.doubanio.com/f/shire/95cc2fa733221bb8edd28ad56a7145a5ad33383e/pics/rating_icons/star_hollow_hover@2x.png" id="star1" width="16" height="16">
        </a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-10594787-2">
            <img src="https://img3.doubanio.com/f/shire/95cc2fa733221bb8edd28ad56a7145a5ad33383e/pics/rating_icons/star_hollow_hover@2x.png" id="star2" width="16" height="16">
        </a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-10594787-3">
            <img src="https://img3.doubanio.com/f/shire/95cc2fa733221bb8edd28ad56a7145a5ad33383e/pics/rating_icons/star_hollow_hover@2x.png" id="star3" width="16" height="16">
        </a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-10594787-4">
            <img src="https://img3.doubanio.com/f/shire/95cc2fa733221bb8edd28ad56a7145a5ad33383e/pics/rating_icons/star_hollow_hover@2x.png" id="star4" width="16" height="16">
        </a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-10594787-5">
            <img src="https://img3.doubanio.com/f/shire/95cc2fa733221bb8edd28ad56a7145a5ad33383e/pics/rating_icons/star_hollow_hover@2x.png" id="star5" width="16" height="16">
        </a>
    </span><span id="rateword" class="pl"></span>
    <input id="n_rating" type="hidden" value="">
    </span>

      </div>
    </div>



  
  <div class="gtleft">
    <ul class="ul_subject_menu bicelink color_gray pt6 clearfix">
        <li>
          <img src="https://img3.doubanio.com/f/shire/5bbf02b7b5ec12b23e214a580b6f9e481108488c/pics/add-review.gif">&nbsp;
          <a class="j a_show_login" href="https://book.douban.com/annotation/write?sid=10594787" rel="nofollow">写笔记</a>
        </li>

          <li>
            <img src="https://img3.doubanio.com/f/shire/5bbf02b7b5ec12b23e214a580b6f9e481108488c/pics/add-review.gif">&nbsp;<a class="j a_show_login" href="https://book.douban.com/subject/10594787/new_review" rel="nofollow">写书评</a>
          </li>

      <li>

<span class="rr">


    <img src="https://img3.doubanio.com/pics/add-cart.gif">
      <a class="j a_show_login" href="https://www.douban.com/accounts/passport/login?source=book" rel="nofollow">加入购书单</a>
  <span class="hidden">已在<a href="https://book.douban.com/cart">购书单</a></span>
</span><br class="clearfix">
</li>


        
        
    
    <li class="rec" id="C-10594787">
        <a href="#" data-url="https://book.douban.com/subject/10594787/" data-desc="" data-title="书籍《霍乱时期的爱情》 (来自豆瓣) " data-pic="https://img2.doubanio.com/view/subject/l/public/s11284102.jpg" class="bn-sharing ">分享到</a> &nbsp;&nbsp;
    </li>
    <script>
    var __cache_url = __cache_url || {};
    (function(u){
        if(__cache_url[u]) return;
        __cache_url[u] = true;
        window.DoubanShareIcons = 'https://img3.doubanio.com/f/shire/d15ffd71f3f10a7210448fec5a68eaec66e7f7d0/pics/ic_shares.png';
        var initShareButton = function() {
          $.ajax({url:u,dataType:'script',cache:true});
        };
        if (typeof Do == 'function' && 'ready' in Do) {
          Do('https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css',
            'https://img3.doubanio.com/f/shire/383a6e43f2108dc69e3ff2681bc4dc6c72a5ffb0/js/ui/dialog.js',
            initShareButton);
        } else if(typeof Douban == 'object' && 'loader' in Douban) {
          Douban.loader.batch(
            'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css',
            'https://img3.doubanio.com/f/shire/383a6e43f2108dc69e3ff2681bc4dc6c72a5ffb0/js/ui/dialog.js'
          ).done(initShareButton);
        }
    })('https://img3.doubanio.com/f/shire/6e6a5f21daeec19bbb41bf48c07fccaa4dad4d98/js/lib/sharebutton.js');
    </script>

    </ul>
  </div>


    







<div class="rec-sec">

    <span class="rec">

<a href="https://www.douban.com/accounts/register?reason=collect" class="j a_show_login lnk-sharing lnk-douban-sharing">推荐</a>
</span>
</div>


<script>
  //bind events for collection button.
  $('.collect_btn', '#interest_sect_level').each(function(){
      Douban.init_collect_btn(this);
  });
</script>








</div>

<br clear="all">
<div id="collect_form_10594787"></div>
<div class="related_info">
  






  

  <h2>
    <span>内容简介</span>
      &nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·

  </h2>



<div class="indent" id="link-report">
    
      <div class="">
        <style type="text/css" media="screen">
.intro p{text-indent:2em;word-break:normal;}
</style>
<div class="intro">
    <p>★马尔克斯唯一正式授权，首次完整翻译</p>    <p>★《霍乱时期的爱情》是我最好的作品，是我发自内心的创作。——加西亚•马尔克斯</p>    <p>★这部光芒闪耀、令人心碎的作品是人类有史以来最伟大的爱情小说。——《纽约时报》</p>    <p>《霍乱时期的爱情》是加西亚•马尔克斯获得诺贝尔文学奖之后完成的第一部小说。讲述了一段跨越半个多世纪的爱情史诗，穷尽了所有爱情的可能性：忠贞的、隐秘的、粗暴的、羞怯的、柏拉图式的、放荡的、转瞬即逝的、生死相依的……再现了时光的无情流逝，被誉为“人类有史以来最伟大的爱情小说”，是20世纪最重要的经典文学巨著之一。</p></div>

      </div>
    
<link rel="stylesheet" type="text/css" href="https://img3.doubanio.com/f/shire/c4c6dd266f58b41cbeebc9e4e6d7dd7b2a5c3711/css/report.css">
<link rel="stylesheet" type="text/css" href="https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css">
<link rel="stylesheet" type="text/css" href="https://img3.doubanio.com/f/shire/b45aa277f8b8df40596b96582dafb1ed0a899a64/css/report_dialog.css">
<script type="text/javascript" src="https://img3.doubanio.com/f/shire/072a17e1a083e85fab3d9bb187de25d07c97f544/js/report_dialog.js"></script>
<style>
    #link-report .report { text-align: right; font-size: 12px; visibility: hidden; }
    #link-report .report a { color: #BBB; }
    #link-report .report a:hover { color: #FFF; background-color: #BBB; }
</style>
<script>
    Do = (typeof Do === 'undefined')? $ : Do;
    Do(function(){
    $("body").delegate("#link-report", 'mouseenter mouseleave', function(e){
      switch (e.type) {
        case "mouseenter":
          $(this).find(".report").css('visibility', 'visible');
          break;
        case "mouseleave":
          $(this).find(".report").css('visibility', 'hidden');
          break;
      }
    });
    $("#link-report").delegate(".report a", 'click', function(e){
        e.preventDefault();
        var opt = "";
        var obj = $(e.target).closest('#link-report');
        var id = obj.length != 0 ? obj.data("id") : undefined;
        var params = (opt&&id) ? ''.concat(opt, '=', id) : '';

        var url = "https://book.douban.com/subject/10594787/";
        url += (~url.indexOf('?') ? '&' : '?') + params
        url = url.replace(/\&+/g, '&')
        generate_report_dialog({report_url: url, type: 'subject'});
    });

    $("#link-report").append('<div class="report"><a rel="nofollow" href="#">投诉</a></div>');
  });
</script><div class="report" style="visibility: hidden;"><a rel="nofollow" href="#">投诉</a></div>

</div>

  

<style type="text/css" media="screen">
  .online-partner{display:flex;align-items:center;padding-top:10px;padding-bottom:10px;margin-bottom:20px;border-bottom:1px solid #D8D8D8}.online-partner .online-type{display:flex;align-items:center}.online-partner .online-type:nth-child(2){padding-left:20px}.online-read-or-audio{display:inline-block;color:rgba(0,0,0,0.9);font-size:12px;line-height:24px}.online-read-or-audio a{display:flex;align-items:center;margin-right:10px;padding:0 15px 0 7px;border:1px solid rgba(0,0,0,0.25);border-radius:5px;vertical-align:middle;color:rgba(0,0,0,0.9)}.online-read-or-audio a:after{content:'';width:5px;height:5px;border-right:1px solid rgba(0,0,0,0.8);border-top:1px solid rgba(0,0,0,0.8);transform:rotate(45deg);margin-right:-6px}.online-read-or-audio a:hover{background:#fff;opacity:.8}.online-read-or-audio img{border-radius:50%;box-shadow:0 0 1px 0 rgba(0,0,0,0.6)}.online-read-or-audio span{padding:0 3px 0 5px}

</style>


  

























    
  

  <h2>
    <span>作者简介</span>
      &nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·

  </h2>



      <div class="indent ">
          
            <div class="">
            <style type="text/css" media="screen">
.intro p{text-indent:2em;word-break:normal;}
</style>
<div class="intro">
    <p>加西亚•马尔克斯（Gabriel García Márquez）</p>    <p>1927年出生于哥伦比亚马格达莱纳海滨小镇阿拉卡塔卡。童年与外祖父母一起生活。1936年随父母迁居苏克雷。1947年考入波哥大国立大学。1948年因内战辍学，进入报界。五十年代开始出版文学作品。六十年代初移居墨西哥。1967年《百年孤独》问世。1982年获诺贝尔文学奖。1985年出版《霍乱时期的爱情》。</p>    <p>加西亚•马尔克斯豆瓣小站：http://site.douban.com/marquez/</p></div>

            </div>
      </div>











































  

  


  


<link rel="stylesheet" href="https://img3.doubanio.com/f/verify/16c7e943aee3b1dc6d65f600fcc0f6d62db7dfb4/entry_creator/dist/author_subject/style.css">

<div id="author_subject" class="author-wrapper"><div data-reactroot="" class="author-subject"></div></div>

<script type="text/javascript">
  var answerObj = {
    TYPE: 'book',
    SUBJECT_ID: '10594787',
    ISALL: 'False' || false,
    USER_ID: 'None'
  }
</script>
<script src="https://img3.doubanio.com/f/book/61252f2f9b35f08b37f69d17dfe48310dd295347/js/book/lib/react/bundle.js"></script>
<script type="text/javascript" src="https://img3.doubanio.com/f/verify/ac140ef86262b845d2be7b859e352d8196f3f6d4/entry_creator/dist/author_subject/index.js"></script> 
  








<div class="ugc-mod blockquote-list-wrapper">
  <div class="hd">
    <h2>
      原文摘录&nbsp;&nbsp;· · · · · ·&nbsp;
      <span class="pl">( <a href="blockquotes">全部</a> )</span>
    </h2>

    <ul class="blockquote-list">
      
        <li>
          <figure>
            当一个女人决定和一个男人睡觉时，就没有她跃不过去的围墙，没有她推不倒的堡垒，也没有她抛不下的道德顾虑，事实上没有能管得住她的上帝。 (<a href="https://book.douban.com/annotation/21176754/">查看原文</a>)

            <div class="blockquote-extra">
              <div class="blockquote-meta">
                <a href="https://www.douban.com/people/56641708/" class="author-avatar">
                  <img width="20" height="20" src="https://img2.doubanio.com/icon/u56641708-2.jpg">
                </a>
                <a class="author-name" href="https://www.douban.com/people/56641708/">慢歌3</a>
                  <span>84 回复</span>
                  <span>696赞</span>
                <datetime>2012-09-18 14:22:01</datetime>
              </div>

              
              <figcaption title="引自第379页">
                —— 引自第379页
              </figcaption>
            </div>
          </figure>
        </li>
      
        <li>
          <figure>
            她因年龄而减损的，又因性格而弥补回来，更因勤劳赢得了更多。 (<a href="https://book.douban.com/annotation/21294558/">查看原文</a>)

            <div class="blockquote-extra">
              <div class="blockquote-meta">
                <a href="https://www.douban.com/people/53240552/" class="author-avatar">
                  <img width="20" height="20" src="https://img9.doubanio.com/icon/u53240552-25.jpg">
                </a>
                <a class="author-name" href="https://www.douban.com/people/53240552/">糖渍柠檬</a>
                  <span>16 回复</span>
                  <span>311赞</span>
                <datetime>2012-09-23 18:01:43</datetime>
              </div>

              
              <figcaption title="引自第28页">
                —— 引自第28页
              </figcaption>
            </div>
          </figure>
        </li>
    </ul>

      <p class="pl"> &gt; <a href="blockquotes">全部原文摘录</a> </p>
  </div>
</div>

  






<div id="db-tags-section" class="blank20">
  
  

  <h2>
    <span>豆瓣成员常用的标签(共3167个)</span>
      &nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·

  </h2>


  <div class="indent">    <span class="">
        <a class="  tag" href="/tag/霍乱时期的爱情">霍乱时期的爱情</a> &nbsp;    </span>
    <span class="">
        <a class="  tag" href="/tag/爱情">爱情</a> &nbsp;    </span>
    <span class="">
        <a class="  tag" href="/tag/加西亚-马尔克斯">加西亚-马尔克斯</a> &nbsp;    </span>
    <span class="">
        <a class="  tag" href="/tag/加西亚·马尔克斯">加西亚·马尔克斯</a> &nbsp;    </span>
    <span class="">
        <a class="  tag" href="/tag/小说">小说</a> &nbsp;    </span>
    <span class="">
        <a class="  tag" href="/tag/拉美文学">拉美文学</a> &nbsp;    </span>
    <span class="">
        <a class="  tag" href="/tag/外国文学">外国文学</a> &nbsp;    </span>
    <span class="">
        <a class="  tag" href="/tag/文学">文学</a> &nbsp;    </span>
  </div>
</div>

  


<div class="subject_show block5">
<h2>丛书信息</h2>
<div>
　　<a href="https://book.douban.com/series/10489">新经典文库：加西亚·马尔克斯作品 (共28册)</a>,
这套丛书还有
《枯枝败叶》,《一桩事先张扬的凶杀案》,《霍乱时期的爱情》,《一起连环绑架案的新闻》,《百年孤独》    等。</div>
</div>
<script>
$(function(){$(".knnlike a").click(function(){return moreurl(this,{'from':'knnlike'})})})
</script>

  












<div id="rec-ebook-section" class="block5 subject_show">
  

  
  

  <h2>
    <span>喜欢读"霍乱时期的爱情"的人也喜欢的电子书</span>
      &nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·

  </h2>


  <div class="tips-mod">
    支持 Web、iPhone、iPad、Android 阅读器
  </div>
  <div class="content clearfix">
      
      <dl>
        <dt>
          <a href="https://read.douban.com/ebook/109905425/?dcs=subject-rec&amp;dcm=douban&amp;dct=10594787" target="_blank">
            <span class="cover-outer">
              <img src="https://img9.doubanio.com/view/ark_article_cover/retina/public/109905425.jpg?v=1553572062">
            </span>
          </a>
        </dt>
        <dd>
          <div class="title">
              <a href="https://read.douban.com/ebook/109905425/" target="_blank">树上的男爵（卡尔维诺经典）</a>
          </div>
          <div class="price">
              
    9.99元

          </div>
        </dd>
      </dl>
      
      <dl>
        <dt>
          <a href="https://read.douban.com/bundle/169371278/?dcs=subject-rec&amp;dcm=douban&amp;dct=10594787" target="_blank">
            <span class="cover-outer">
              <img src="https://img1.doubanio.com/view/ark_article_cover/retina/public/169371278.jpg?v=1610693803">
            </span>
          </a>
        </dt>
        <dd>
          <div class="title">
              <a href="https://read.douban.com/bundle/169371278/" target="_blank">红楼梦（套装全2册）</a>
          </div>
          <div class="price">
              
    36.00元

          </div>
        </dd>
      </dl>
      
      <dl>
        <dt>
          <a href="https://read.douban.com/ebook/754089/?dcs=subject-rec&amp;dcm=douban&amp;dct=10594787" target="_blank">
            <span class="cover-outer">
              <img src="https://img1.doubanio.com/view/ark_article_cover/retina/public/754089.jpg?v=0">
            </span>
          </a>
        </dt>
        <dd>
          <div class="title">
              <a href="https://read.douban.com/ebook/754089/" target="_blank">局外人</a>
          </div>
          <div class="price">
              
    0.99元

          </div>
        </dd>
      </dl>
      
      <dl>
        <dt>
          <a href="https://read.douban.com/ebook/407582/?dcs=subject-rec&amp;dcm=douban&amp;dct=10594787" target="_blank">
            <span class="cover-outer">
              <img src="https://img2.doubanio.com/view/ark_article_cover/retina/public/407582.jpg?v=0">
            </span>
          </a>
        </dt>
        <dd>
          <div class="title">
              <a href="https://read.douban.com/ebook/407582/" target="_blank">少年PI的奇幻漂流（插图珍藏版）</a>
          </div>
          <div class="price">
              
    5.99元

          </div>
        </dd>
      </dl>
      
      <dl>
        <dt>
          <a href="https://read.douban.com/ebook/316702/?dcs=subject-rec&amp;dcm=douban&amp;dct=10594787" target="_blank">
            <span class="cover-outer">
              <img src="https://img2.doubanio.com/view/ark_article_cover/retina/public/316702.jpg?v=0">
            </span>
          </a>
        </dt>
        <dd>
          <div class="title">
              <a href="https://read.douban.com/ebook/316702/" target="_blank">论文艺女青年如何培养女王气场</a>
          </div>
          <div class="price">
              
    1.99元

          </div>
        </dd>
      </dl>
  </div>
</div>

<div id="db-rec-section" class="block5 subject_show knnlike">
  
  
  

  <h2>
    <span>喜欢读"霍乱时期的爱情"的人也喜欢</span>
      &nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·

  </h2>


  <div class="content clearfix">
      
      <dl class="">
        <dt>
            <a href="https://book.douban.com/subject/6082808/" onclick="moreurl(this, {'total': 10, 'clicked': '6082808', 'pos': 0, 'identifier': 'book-rec-books'})"><img class="m_sub_img" src="https://img3.doubanio.com/view/subject/s/public/s27237850.jpg"></a>
        </dt>
        <dd>
          <a href="https://book.douban.com/subject/6082808/" onclick="moreurl(this, {'total': 10, 'clicked': '6082808', 'pos': 0, 'identifier': 'book-rec-books'})" class="">
            百年孤独
          </a>
        </dd>
      </dl>
      
      <dl class="">
        <dt>
            <a href="https://book.douban.com/subject/1083428/" onclick="moreurl(this, {'total': 10, 'clicked': '1083428', 'pos': 1, 'identifier': 'book-rec-books'})"><img class="m_sub_img" src="https://img2.doubanio.com/view/subject/s/public/s4250062.jpg"></a>
        </dt>
        <dd>
          <a href="https://book.douban.com/subject/1083428/" onclick="moreurl(this, {'total': 10, 'clicked': '1083428', 'pos': 1, 'identifier': 'book-rec-books'})" class="">
            傲慢与偏见
          </a>
        </dd>
      </dl>
      
      <dl class="">
        <dt>
            <a href="https://book.douban.com/subject/4820710/" onclick="moreurl(this, {'total': 10, 'clicked': '4820710', 'pos': 2, 'identifier': 'book-rec-books'})"><img class="m_sub_img" src="https://img1.doubanio.com/view/subject/s/public/s4371408.jpg"></a>
        </dt>
        <dd>
          <a href="https://book.douban.com/subject/4820710/" onclick="moreurl(this, {'total': 10, 'clicked': '4820710', 'pos': 2, 'identifier': 'book-rec-books'})" class="">
            1984
          </a>
        </dd>
      </dl>
      
      <dl class="">
        <dt>
            <a href="https://book.douban.com/subject/1084336/" onclick="moreurl(this, {'total': 10, 'clicked': '1084336', 'pos': 3, 'identifier': 'book-rec-books'})"><img class="m_sub_img" src="https://img2.doubanio.com/view/subject/s/public/s1103152.jpg"></a>
        </dt>
        <dd>
          <a href="https://book.douban.com/subject/1084336/" onclick="moreurl(this, {'total': 10, 'clicked': '1084336', 'pos': 3, 'identifier': 'book-rec-books'})" class="">
            小王子
          </a>
        </dd>
      </dl>
      
      <dl class="">
        <dt>
            <a href="https://book.douban.com/subject/1008988/" onclick="moreurl(this, {'total': 10, 'clicked': '1008988', 'pos': 4, 'identifier': 'book-rec-books'})"><img class="m_sub_img" src="https://img9.doubanio.com/view/subject/s/public/s1005875.jpg"></a>
        </dt>
        <dd>
          <a href="https://book.douban.com/subject/1008988/" onclick="moreurl(this, {'total': 10, 'clicked': '1008988', 'pos': 4, 'identifier': 'book-rec-books'})" class="">
            了不起的盖茨比
          </a>
        </dd>
      </dl>
        <dl class="clear"></dl>
      
      <dl class="">
        <dt>
            <a href="https://book.douban.com/subject/4011670/" onclick="moreurl(this, {'total': 10, 'clicked': '4011670', 'pos': 5, 'identifier': 'book-rec-books'})"><img class="m_sub_img" src="https://img9.doubanio.com/view/subject/s/public/s6100756.jpg"></a>
        </dt>
        <dd>
          <a href="https://book.douban.com/subject/4011670/" onclick="moreurl(this, {'total': 10, 'clicked': '4011670', 'pos': 5, 'identifier': 'book-rec-books'})" class="">
            人间失格
          </a>
        </dd>
      </dl>
      
      <dl class="">
        <dt>
            <a href="https://book.douban.com/subject/1200840/" onclick="moreurl(this, {'total': 10, 'clicked': '1200840', 'pos': 6, 'identifier': 'book-rec-books'})"><img class="m_sub_img" src="https://img3.doubanio.com/view/subject/s/public/s1144911.jpg"></a>
        </dt>
        <dd>
          <a href="https://book.douban.com/subject/1200840/" onclick="moreurl(this, {'total': 10, 'clicked': '1200840', 'pos': 6, 'identifier': 'book-rec-books'})" class="">
            平凡的世界（全三部）
          </a>
        </dd>
      </dl>
      
      <dl class="">
        <dt>
            <a href="https://book.douban.com/subject/1068920/" onclick="moreurl(this, {'total': 10, 'clicked': '1068920', 'pos': 7, 'identifier': 'book-rec-books'})"><img class="m_sub_img" src="https://img1.doubanio.com/view/subject/s/public/s1078958.jpg"></a>
        </dt>
        <dd>
          <a href="https://book.douban.com/subject/1068920/" onclick="moreurl(this, {'total': 10, 'clicked': '1068920', 'pos': 7, 'identifier': 'book-rec-books'})" class="">
            飘
          </a>
        </dd>
      </dl>
      
      <dl class="">
        <dt>
            <a href="https://book.douban.com/subject/2154960/" onclick="moreurl(this, {'total': 10, 'clicked': '2154960', 'pos': 8, 'identifier': 'book-rec-books'})"><img class="m_sub_img" src="https://img1.doubanio.com/view/subject/s/public/s2611329.jpg"></a>
        </dt>
        <dd>
          <a href="https://book.douban.com/subject/2154960/" onclick="moreurl(this, {'total': 10, 'clicked': '2154960', 'pos': 8, 'identifier': 'book-rec-books'})" class="">
            一个陌生女人的来信
          </a>
        </dd>
      </dl>
      
      <dl class="">
        <dt>
            <a href="https://book.douban.com/subject/1141406/" onclick="moreurl(this, {'total': 10, 'clicked': '1141406', 'pos': 9, 'identifier': 'book-rec-books'})"><img class="m_sub_img" src="https://img9.doubanio.com/view/subject/s/public/s5924326.jpg"></a>
        </dt>
        <dd>
          <a href="https://book.douban.com/subject/1141406/" onclick="moreurl(this, {'total': 10, 'clicked': '1141406', 'pos': 9, 'identifier': 'book-rec-books'})" class="">
            简爱（英文全本）
          </a>
        </dd>
      </dl>
        <dl class="clear"></dl>
  </div>
</div>

  





    <div id="comments-section">
        <link rel="stylesheet" href="https://img3.doubanio.com/f/book/16c2b96ce936f65269d562970bdfd7b3572ffec9/css/book/lib/subject-comments/comments-section.css">
        <div class="mod-hd">
            

            <a class="redbutt j a_show_login rr" href="https://www.douban.com/register?reason=review" rel="nofollow">
                <span> 我来说两句 </span>
            </a>

                
  

  <h2>
    <span>短评</span>
      &nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·
      <span class="pl">&nbsp;(
          <a href="https://book.douban.com/subject/10594787/comments/">全部 62074 条</a>
        ) </span>

  </h2>


        </div>
        <div class="nav-tab">
            
    <div class="tabs-wrapper  line">
        <a class="short-comment-tabs on-tab" href="https://book.douban.com/subject/10594787/comments?sort=new_score" data-tab="new_score">热门</a>
        <span>/</span>
        <a class="short-comment-tabs " href="https://book.douban.com/subject/10594787/comments?sort=time" data-tab="time">最新</a>
        <span>/</span>
        <a class="j a_show_login " href="https://book.douban.com/subject/10594787/comments?sort=follows" data-tab="follows">好友</a>
    </div>

        </div>
        <div id="comment-list-wrapper" class="indent">
            
  
  <div class="comment-list new_score show" id="new_score">
      <ul>
          
  <li class="comment-item" data-cid="810532324">
    <div class="comment">
      <h3>
        <span class="comment-vote">
          <span id="c-810532324" class="vote-count">1989</span>
            <a href="javascript:;" id="btn-810532324" class="j a_show_login" data-cid="810532324">有用</a>
        </span>
        <span class="comment-info">
          <a href="https://www.douban.com/people/59291955/">丹朱AI🌵</a>
            <span class="user-stars allstar50 rating" title="力荐"></span>
          <span class="comment-time">2014-05-26</span>
        </span>
      </h3>
      <p class="comment-content">
      
        <span class="short">我猜这本四百页的书并非是为了歌颂爱情，而是为了呈现人生，爱情是杂乱而繁茂的生命之树上温情脉脉的一缕夕阳。</span>
      </p>
    </div>
  </li>

          
  <li class="comment-item" data-cid="582083012">
    <div class="comment">
      <h3>
        <span class="comment-vote">
          <span id="c-582083012" class="vote-count">470</span>
            <a href="javascript:;" id="btn-582083012" class="j a_show_login" data-cid="582083012">有用</a>
        </span>
        <span class="comment-info">
          <a href="https://www.douban.com/people/chengzhimo/">成知默</a>
            <span class="user-stars allstar50 rating" title="力荐"></span>
          <span class="comment-time">2012-09-20</span>
        </span>
      </h3>
      <p class="comment-content">
      
        <span class="short">对于爱情，我们都太年轻了，而且我们永远太年轻，不知道人活在这个世界上，爱情有多少面目。</span>
      </p>
    </div>
  </li>

          
  <li class="comment-item" data-cid="894685733">
    <div class="comment">
      <h3>
        <span class="comment-vote">
          <span id="c-894685733" class="vote-count">1942</span>
            <a href="javascript:;" id="btn-894685733" class="j a_show_login" data-cid="894685733">有用</a>
        </span>
        <span class="comment-info">
          <a href="https://www.douban.com/people/wusequanwei/">五色全味</a>
            <span class="user-stars allstar50 rating" title="力荐"></span>
          <span class="comment-time">2015-02-18</span>
        </span>
      </h3>
      <p class="comment-content">
      
        <span class="short">竟然舍不得读完。但也不觉得写的是所谓的“爱情”，哪怕结尾多么大无畏和浪漫。写的其实是“孤独”，无论那女人多么“高傲”，那男人多么“卑微”，她在世俗的婚姻中以为拥有了爱情，他在浪荡滥交的生活中以为逃避了空虚，当他们白发苍苍时相拥在一起，才发现他们这一辈子恒久的寂寞的灵魂</span>
      </p>
    </div>
  </li>

          
  <li class="comment-item" data-cid="685383493">
    <div class="comment">
      <h3>
        <span class="comment-vote">
          <span id="c-685383493" class="vote-count">366</span>
            <a href="javascript:;" id="btn-685383493" class="j a_show_login" data-cid="685383493">有用</a>
        </span>
        <span class="comment-info">
          <a href="https://www.douban.com/people/3868526/">偏见小姐周婉粥</a>
            <span class="user-stars allstar40 rating" title="推荐"></span>
          <span class="comment-time">2013-07-21</span>
        </span>
      </h3>
      <p class="comment-content">
      
        <span class="short">又一个盖茨比的故事？我才不信爱情呢，这种持续五十几年的所谓爱情不过只是偏执而已。比起这种幻觉一样的爱情，我倒觉得乌尔比诺医生那种安放在日常琐碎中的情感更值得人尊重。另外，得不到一个女人，就要靠从此放浪形骸来平衡内心的失落吗？人呐，没事少装纳兰公子，假借爱情之名恣意纵情，伤及无辜的那些人，我看原本就是些不配得到爱情的。</span>
      </p>
    </div>
  </li>

          
  <li class="comment-item" data-cid="581444921">
    <div class="comment">
      <h3>
        <span class="comment-vote">
          <span id="c-581444921" class="vote-count">6454</span>
            <a href="javascript:;" id="btn-581444921" class="j a_show_login" data-cid="581444921">有用</a>
        </span>
        <span class="comment-info">
          <a href="https://www.douban.com/people/Francyslio/">飞行官</a>
            <span class="user-stars allstar50 rating" title="力荐"></span>
          <span class="comment-time">2012-09-18</span>
        </span>
      </h3>
      <p class="comment-content">
      
        <span class="short">不夸张地讲，这本书包含了爱情的全部答案。阅读本书的好处是重新相信爱情；坏处是意识到即便相信也无济于事。</span>
      </p>
    </div>
  </li>

      </ul>
  </div>
  


            
  
  <div class="comment-list time noshow" id="time">
      <ul>
          
  <li class="comment-item" data-cid="2800721982">
    <div class="comment">
      <h3>
        <span class="comment-vote">
          <span id="c-2800721982" class="vote-count">0</span>
            <a href="javascript:;" id="btn-2800721982" class="j a_show_login" data-cid="2800721982">有用</a>
        </span>
        <span class="comment-info">
          <a href="https://www.douban.com/people/202971383/">Hikikomori</a>
            <span class="user-stars allstar50 rating" title="力荐"></span>
          <span class="comment-time">2021-03-21</span>
        </span>
      </h3>
      <p class="comment-content">
      
        <span class="short">最喜欢的爱情小说，寒假的下午温暖的阳光和这本书，欲罢不能荡气回肠</span>
      </p>
    </div>
  </li>

          
  <li class="comment-item" data-cid="2800697872">
    <div class="comment">
      <h3>
        <span class="comment-vote">
          <span id="c-2800697872" class="vote-count">0</span>
            <a href="javascript:;" id="btn-2800697872" class="j a_show_login" data-cid="2800697872">有用</a>
        </span>
        <span class="comment-info">
          <a href="https://www.douban.com/people/139886347/">安小稚</a>
            <span class="user-stars allstar50 rating" title="力荐"></span>
          <span class="comment-time">2021-03-21</span>
        </span>
      </h3>
      <p class="comment-content">
      
        <span class="short">安全感、和谐和幸福，这些东西一旦相加，或许看似爱情，也几乎等于爱情，但他们终究不是爱情。</span>
      </p>
    </div>
  </li>

          
  <li class="comment-item" data-cid="1022725204">
    <div class="comment">
      <h3>
        <span class="comment-vote">
          <span id="c-1022725204" class="vote-count">0</span>
            <a href="javascript:;" id="btn-1022725204" class="j a_show_login" data-cid="1022725204">有用</a>
        </span>
        <span class="comment-info">
          <a href="https://www.douban.com/people/139384223/">柏林</a>
            <span class="user-stars allstar50 rating" title="力荐"></span>
          <span class="comment-time">2021-03-21</span>
        </span>
      </h3>
      <p class="comment-content">
      
        <span class="short">2021.03.21
花了一个月看完了五年前想看的《霍乱时期的爱情》
与原本想象中的霍乱不同，我认为书名的霍乱最主要体现在人心里，爱情就好像霍乱一样扎根在人们心中。
在跨越了半个多世纪，53年7个月零11天后，佛洛伦蒂诺·阿里萨与费尔明娜·达萨终于在二人晚年的日子里拥抱了自己的爱情。尽管在这半个多世纪中二人似乎都有着各自的生活，但现在再也没有任何人事物能拆散彼此了。</span>
      </p>
    </div>
  </li>

          
  <li class="comment-item" data-cid="2534429075">
    <div class="comment">
      <h3>
        <span class="comment-vote">
          <span id="c-2534429075" class="vote-count">0</span>
            <a href="javascript:;" id="btn-2534429075" class="j a_show_login" data-cid="2534429075">有用</a>
        </span>
        <span class="comment-info">
          <a href="https://www.douban.com/people/158803355/">于年鱼</a>
            <span class="user-stars allstar50 rating" title="力荐"></span>
          <span class="comment-time">2021-03-21</span>
        </span>
      </h3>
      <p class="comment-content">
      
        <span class="short">看完了我依然不懂爱情，但还是被这里的爱情震惊了</span>
      </p>
    </div>
  </li>

          
  <li class="comment-item" data-cid="2800530865">
    <div class="comment">
      <h3>
        <span class="comment-vote">
          <span id="c-2800530865" class="vote-count">0</span>
            <a href="javascript:;" id="btn-2800530865" class="j a_show_login" data-cid="2800530865">有用</a>
        </span>
        <span class="comment-info">
          <a href="https://www.douban.com/people/208267247/">齐齐</a>
            <span class="user-stars allstar50 rating" title="力荐"></span>
          <span class="comment-time">2021-03-21</span>
        </span>
      </h3>
      <p class="comment-content">
      
        <span class="short">马尔克斯的文笔太好了</span>
      </p>
    </div>
  </li>

      </ul>
  </div>
  


        </div>
            <p>&gt; <a href="https://book.douban.com/subject/10594787/comments/">更多短评 62074 条</a></p>
        <script src="https://img3.doubanio.com/f/book/6eba6f43fb7592ab783e390f654c0d6a96b1598e/js/book/lib/subject-comments/comments-section.js"></script>
        <script>
            (function () {
                if (window.SUBJECT_COMMENTS_SECTION) {
                    // tab handler
                    SUBJECT_COMMENTS_SECTION.createTabHandler();
                    // expand handler
                    SUBJECT_COMMENTS_SECTION.createExpandHandler({
                        root: document.getElementById('comment-list-wrapper'),
                    });
                    SUBJECT_COMMENTS_SECTION.createVoteHandler({
                        api: '/j/comment/:id/vote',
                        root: document.getElementById('comment-list-wrapper'),
                        voteSelector: '.vote-comment',
                        textSelector: '.vote-count',
                        afterVote: function (elem) {
                            var parentNode = elem.parentNode;
                            var successElem = document.createElement('span');
                            successElem.innerHTML = '已投票';
                            parentNode.removeChild(elem);
                            parentNode.appendChild(successElem);
                        }
                    });
                }
            })()
        </script>
    </div>




  

<link rel="stylesheet" href="https://img3.doubanio.com/misc/mixed_static/1b80cf3bee515ce2.css">


    <section id="reviews-wrapper" class="reviews mod book-content">
        <header>
            
                <a href="new_review" rel="nofollow" class="create-review redbutt rr " data-isverify="False" data-verify-url="https://www.douban.com/accounts/phone/verify?redir=https://book.douban.com/subject/10594787/new_review">
                    <span>我要写书评</span>
                </a>
            <h2>
                霍乱时期的爱情的书评 · · · · · ·
                <span class="pl">( <a href="reviews">全部 4005 条</a> )</span>
            </h2>
        </header>

        

<style>
#gallery-topics-selection {
  position: fixed;
  width: 595px;
  padding: 40px 40px 33px 40px;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 16px 0 rgba(0, 0, 0, 0.2);
  top: 50%;
  left: 50%;
  -webkit-transform: translate(-50%, -50%);
  transform: translate(-50%, -50%);
  z-index: 9999;
}
#gallery-topics-selection h1 {
  font-size: 18px;
  color: #007722;
  margin-bottom: 36px;
  padding: 0;
  line-height: 28px;
  font-weight: normal;
}
#gallery-topics-selection .gl_topics {
  border-bottom: 1px solid #dfdfdf;
  max-height: 298px;
  overflow-y: scroll;
}
#gallery-topics-selection .topic {
  margin-bottom: 24px;
}
#gallery-topics-selection .topic_name {
  font-size: 15px;
  color: #333;
  margin: 0;
  line-height: inherit;
}
#gallery-topics-selection .topic_meta {
  font-size: 13px;
  color: #999;
}
#gallery-topics-selection .topics_skip {
  display: block;
  cursor: pointer;
  font-size: 16px;
  color: #3377AA;
  text-align: center;
  margin-top: 33px;
}
#gallery-topics-selection .topics_skip:hover {
  background: transparent;
}
#gallery-topics-selection .close_selection {
  position: absolute;
  width: 30px;
  height: 20px;
  top: 46px;
  right: 40px;
  background: #fff;
  color: #999;
  text-align: right;
}
#gallery-topics-selection .close_selection:hover{
  background: #fff;
  color: #999;
}
</style>




            <div class="review_filter">
                <a href="javascript:;;" class="cur" data-sort="">热门</a> /
                <a href="javascript:;;" data-sort="time">最新</a> /
                <a href="javascript:;;" data-sort="follow">好友</a>
                    /<a href="/subject/10594787/reviews?version=1" data-sort="vertion">只看本版本的评论</a>
            </div>
            <script>
                var cur_sort = '';
                $('#reviews-wrapper .review_filter a').on('click', function () {
                    var sort = $(this).data('sort');
                    if(sort === cur_sort) return;

                    if(sort === 'follow' && true){
                        window.location.href = '//www.douban.com/accounts/login?source=movie';
                        return;
                    }

                    if($('#reviews-wrapper .review_filter').data('doing')) return;
                    $('#reviews-wrapper .review_filter').data('doing', true);

                    cur_sort = sort;

                    $('#reviews-wrapper .review_filter a').removeClass('cur');
                    $(this).addClass('cur');

                    $.getJSON('reviews', { sort: sort }, function(res){
                        $('#reviews-wrapper .review-list').remove();
                        $('#reviews-wrapper [href="reviews?sort=follow"]').parent().remove();
                        $('#reviews-wrapper .review_filter').after(res.html);
                        $('#reviews-wrapper .review_filter').data('doing', false);
                        $('#reviews-wrapper .review_filter').removeData('doing');

                        if(res.count === 0){
                            $('#reviews-wrapper .review-list').html('<span class="no-review">你关注的人还没写过长评</span>');
                        }
                    });
                });
            </script>


            



<div class="review-list  ">
        
    

        
    
    <div data-cid="6706506">
        <div class="main review-item" id="6706506">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/nrl/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u3060807-250.jpg">
        </a>

        <a href="https://www.douban.com/people/nrl/" class="name">吧啦</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2014-06-18" class="main-meta">2014-06-18 21:05:48</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://book.douban.com/review/6706506/">马尔克斯的玩笑</a></h2>

                <div id="review_6706506_short" class="review-short" data-rid="6706506">
                    <div class="short-content">

                        “一生一世。”   合上《霍乱时期的爱情》这本书时，我仿佛看到马尔克斯在角落里偷笑。  两年前读完这本书，就觉得有哪里不对。再读一次，彻底感觉被马尔克斯耍了——这本所谓的“爱情的百科全书”，哪里是在写爱情呀。  一个叫佛洛伦蒂诺•阿里萨的文艺青年，看上了一个叫费...

                        &nbsp;(<a href="javascript:;" id="toggle-6706506-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_6706506_full" class="hidden">
                    <div id="review_6706506_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="6706506" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png">
                        <span id="r-useful_count-6706506">
                                5081
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="6706506" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png">
                        <span id="r-useless_count-6706506">
                                315
                        </span>
                    </a>
                    <a href="https://book.douban.com/review/6706506/#comments" class="reply ">692回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="1547774">
        <div class="main review-item" id="1547774">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/renxiaowen/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u1176907-21.jpg">
        </a>

        <a href="https://www.douban.com/people/renxiaowen/" class="name">任晓雯</a>

            <span class="allstar50 main-title-rating" title="力荐"></span>

        <span content="2008-11-06" class="main-meta">2008-11-06 19:41:47</span>


            <span class="publisher right">
                <a class="publisher" target="_blank" href="https://book.douban.com/subject/1292227/">敦煌文艺出版社2000版</a>
            </span>
    </header>


            <div class="main-bd">

                <h2><a href="https://book.douban.com/review/1547774/">打字机情书与暮年的白玫瑰——读《霍乱时期的爱情》</a></h2>

                <div id="review_1547774_short" class="review-short" data-rid="1547774">
                    <div class="short-content">

                        一  如果在阅读中掩去作者的姓名背景，我也将毫不怀疑地断定，这本充满迟暮感伤的书，出自一位老者。不过在此之前，我已获得了关于此书的初步印象：它完成于1985年，当时57岁的加西亚•马尔克斯，于四年前得到诺贝尔奖，正享有着与日俱增的世界性荣耀。 作为无愧于“大师”称...

                        &nbsp;(<a href="javascript:;" id="toggle-1547774-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_1547774_full" class="hidden">
                    <div id="review_1547774_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="1547774" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png">
                        <span id="r-useful_count-1547774">
                                2815
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="1547774" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png">
                        <span id="r-useless_count-1547774">
                                104
                        </span>
                    </a>
                    <a href="https://book.douban.com/review/1547774/#comments" class="reply ">166回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="5583672">
        <div class="main review-item" id="5583672">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/ofcandy/" class="avator">
            <img width="24" height="24" src="https://img1.doubanio.com/icon/u1086983-337.jpg">
        </a>

        <a href="https://www.douban.com/people/ofcandy/" class="name">糖罐子</a>

            <span class="allstar50 main-title-rating" title="力荐"></span>

        <span content="2012-09-14" class="main-meta">2012-09-14 13:32:25</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://book.douban.com/review/5583672/">我们可以老了再相爱</a></h2>

                <div id="review_5583672_short" class="review-short" data-rid="5583672">
                    <div class="short-content">

                        “我们可以老了再相爱”,我忘了最初是从那听到的这句话了,只记得当时我正反复回味阿飞姑娘念叨的”再不相爱就老了”，总觉得一切的一切都已经来不及了,然后在那个当口,我听到有人说:我们可以老了再相爱。那感觉忽然很棒，对时间的焦灼感好像如释重负。  差不多读到《霍乱时期的...

                        &nbsp;(<a href="javascript:;" id="toggle-5583672-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_5583672_full" class="hidden">
                    <div id="review_5583672_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="5583672" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png">
                        <span id="r-useful_count-5583672">
                                1955
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="5583672" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png">
                        <span id="r-useless_count-5583672">
                                122
                        </span>
                    </a>
                    <a href="https://book.douban.com/review/5583672/#comments" class="reply ">140回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="5574953">
        <div class="main review-item" id="5574953">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/lydy/" class="avator">
            <img width="24" height="24" src="https://img1.doubanio.com/icon/u1364551-9.jpg">
        </a>

        <a href="https://www.douban.com/people/lydy/" class="name">当归莳子</a>

            <span class="allstar50 main-title-rating" title="力荐"></span>

        <span content="2012-09-06" class="main-meta">2012-09-06 12:22:38</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://book.douban.com/review/5574953/">关于新版。兼说翻译。</a></h2>

                <div id="review_5574953_short" class="review-short" data-rid="5574953">
                    <div class="short-content">

                        人世间的翻译，从来没有什么“最高”。那些无论自吹还是他吹所谓史上最牛叉者，无论林老师还是李老师，纯是瞎扯。 额看来，真正的好翻译，很简单，大体就是你既然能翻过来，当然必须还能够还（读huan)回去。 翻过来，新语言能将书里的味和趣能讲的如原书；又还能翻回去，同原文...

                        &nbsp;(<a href="javascript:;" id="toggle-5574953-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_5574953_full" class="hidden">
                    <div id="review_5574953_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="5574953" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png">
                        <span id="r-useful_count-5574953">
                                1423
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="5574953" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png">
                        <span id="r-useless_count-5574953">
                                60
                        </span>
                    </a>
                    <a href="https://book.douban.com/review/5574953/#comments" class="reply ">169回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="6250048">
        <div class="main review-item" id="6250048">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/danyboy/" class="avator">
            <img width="24" height="24" src="https://img1.doubanio.com/icon/u1186285-18.jpg">
        </a>

        <a href="https://www.douban.com/people/danyboy/" class="name">danyboy</a>

            <span class="allstar50 main-title-rating" title="力荐"></span>

        <span content="2013-08-27" class="main-meta">2013-08-27 22:49:29</span>


            <span class="publisher right">
                <a class="publisher" target="_blank" href="https://book.douban.com/subject/1870320/">漓江出版社1987版</a>
            </span>
    </header>


            <div class="main-bd">

                <h2><a href="https://book.douban.com/review/6250048/">三个关于爱情的问题</a></h2>

                <div id="review_6250048_short" class="review-short" data-rid="6250048">
                    <div class="short-content">

                        　　我读的这个漓江出版社1987年由徐鹤林、魏民翻译的版本，是多年前从旧书市场购买的，我觉得翻译的很好。  ●费尔明娜•达萨老去之后爱的是哪个弗洛伦蒂诺？  　　老人的爱情一点也不稀奇，甚至也不令人动心。一辈子守着慢慢变老也好，老去之后重寻伴侣也好，都是平常事。...

                        &nbsp;(<a href="javascript:;" id="toggle-6250048-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_6250048_full" class="hidden">
                    <div id="review_6250048_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="6250048" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png">
                        <span id="r-useful_count-6250048">
                                1148
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="6250048" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png">
                        <span id="r-useless_count-6250048">
                                40
                        </span>
                    </a>
                    <a href="https://book.douban.com/review/6250048/#comments" class="reply ">122回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="1396640">
        <div class="main review-item" id="1396640">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/2546734/" class="avator">
            <img width="24" height="24" src="https://img1.doubanio.com/icon/user_normal.jpg">
        </a>

        <a href="https://www.douban.com/people/2546734/" class="name">吹埙女</a>

            <span class="allstar50 main-title-rating" title="力荐"></span>

        <span content="2008-06-02" class="main-meta">2008-06-02 17:22:50</span>


            <span class="publisher right">
                <a class="publisher" target="_blank" href="https://book.douban.com/subject/2216106/">黑龙江人民出版社1988版</a>
            </span>
    </header>


            <div class="main-bd">

                <h2><a href="https://book.douban.com/review/1396640/">世上的盐，世上的光 ------ 《霍乱时期的爱情》</a></h2>

                <div id="review_1396640_short" class="review-short" data-rid="1396640">
                    <div class="short-content">

                        我不读书已经多年。前些天，一个朋友，我们时代最好最出色的写字者之一，向我推荐马尔克斯的小说《霍乱时期的爱情》。以前曾经读过《百年孤独》，被马尔克斯魔幻得虽然震撼但是迷糊。本以为《霍乱》也是一路货色。可是，据说，在《霍乱》中“他果断放弃了“魔幻现实”的拿手好...

                        &nbsp;(<a href="javascript:;" id="toggle-1396640-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_1396640_full" class="hidden">
                    <div id="review_1396640_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="1396640" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png">
                        <span id="r-useful_count-1396640">
                                593
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="1396640" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png">
                        <span id="r-useless_count-1396640">
                                47
                        </span>
                    </a>
                    <a href="https://book.douban.com/review/1396640/#comments" class="reply ">78回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="5571885">
        <div class="main review-item" id="5571885">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/4646789/" class="avator">
            <img width="24" height="24" src="https://img9.doubanio.com/icon/u4646789-4.jpg">
        </a>

        <a href="https://www.douban.com/people/4646789/" class="name">夕西然</a>

            <span class="allstar50 main-title-rating" title="力荐"></span>

        <span content="2012-09-03" class="main-meta">2012-09-03 20:07:57</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://book.douban.com/review/5571885/">爱在孤独的尽头</a></h2>

                <div id="review_5571885_short" class="review-short" data-rid="5571885">
                    <div class="short-content">

                                前些天，在拥挤、喧嚣的火车上读完了这本关于爱情的不朽之作。当时，在漫长的旅途之中体味一场更为漫长的人生之旅，只觉的沧海桑田，其中细味，还未来得及深深咀嚼。  　　         只是昨天的某些瞬间，走在这平常城市的街道之上，看到油漆门窗的小店门口，一片神秘的...

                        &nbsp;(<a href="javascript:;" id="toggle-5571885-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_5571885_full" class="hidden">
                    <div id="review_5571885_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="5571885" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png">
                        <span id="r-useful_count-5571885">
                                296
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="5571885" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png">
                        <span id="r-useless_count-5571885">
                                21
                        </span>
                    </a>
                    <a href="https://book.douban.com/review/5571885/#comments" class="reply ">39回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="5551268">
        <div class="main review-item" id="5551268">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/52411351/" class="avator">
            <img width="24" height="24" src="https://img1.doubanio.com/icon/u52411351-18.jpg">
        </a>

        <a href="https://www.douban.com/people/52411351/" class="name">enson</a>

            <span class="allstar50 main-title-rating" title="力荐"></span>

        <span content="2012-08-18" class="main-meta">2012-08-18 11:08:55</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://book.douban.com/review/5551268/">加西亚·马尔克斯谈《霍乱时期的爱情》 /申宝楼译</a></h2>

                <div id="review_5551268_short" class="review-short" data-rid="5551268">
                    <div class="short-content">

                        　　问：《霍乱时期的爱情》是你获得诺贝尔文学奖后的第一部小说，获奖一事对你有什么影响？ 　　答：第一个影响是我不得不中断我写作中的小说。得知获奖的消息时，我已经开始写这本书了。当时我曾不切实际地想，举行发奖仪式和接受访者采访后，我便可以重新开始这本小说的写作...

                        &nbsp;(<a href="javascript:;" id="toggle-5551268-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_5551268_full" class="hidden">
                    <div id="review_5551268_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="5551268" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png">
                        <span id="r-useful_count-5551268">
                                260
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="5551268" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png">
                        <span id="r-useless_count-5551268">
                                7
                        </span>
                    </a>
                    <a href="https://book.douban.com/review/5551268/#comments" class="reply ">18回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="5563379">
        <div class="main review-item" id="5563379">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/lydy/" class="avator">
            <img width="24" height="24" src="https://img1.doubanio.com/icon/u1364551-9.jpg">
        </a>

        <a href="https://www.douban.com/people/lydy/" class="name">当归莳子</a>

            <span class="allstar50 main-title-rating" title="力荐"></span>

        <span content="2012-08-28" class="main-meta">2012-08-28 10:44:26</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://book.douban.com/review/5563379/">爱情，蒙蔽了20亿人的眼</a></h2>

                <div id="review_5563379_short" class="review-short" data-rid="5563379">
                    <div class="short-content">

                        据欧洲圈子里的说法，全球完全彻底读过《霍乱时期的爱情》的人超过20亿。 几乎在所有关于《霍乱时期的爱情》的评语中，都有有“爱情的百科全书”字样，天朝一伙伙著名作家\书评人\自媒体也都未能免俗。 确实俗。 所以说俗，是因为那个数字，那个622。 对老马稍微熟悉的一点的人...

                        &nbsp;(<a href="javascript:;" id="toggle-5563379-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_5563379_full" class="hidden">
                    <div id="review_5563379_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="5563379" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png">
                        <span id="r-useful_count-5563379">
                                282
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="5563379" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png">
                        <span id="r-useless_count-5563379">
                                46
                        </span>
                    </a>
                    <a href="https://book.douban.com/review/5563379/#comments" class="reply ">48回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="1029715">
        <div class="main review-item" id="1029715">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/clockcrow/" class="avator">
            <img width="24" height="24" src="https://img2.doubanio.com/icon/u1109069-2.jpg">
        </a>

        <a href="https://www.douban.com/people/clockcrow/" class="name">克劳克罗</a>

            <span class="allstar50 main-title-rating" title="力荐"></span>

        <span content="2006-03-06" class="main-meta">2006-03-06 22:35:27</span>


            <span class="publisher right">
                <a class="publisher" target="_blank" href="https://book.douban.com/subject/1292227/">敦煌文艺出版社2000版</a>
            </span>
    </header>


            <div class="main-bd">

                <h2><a href="https://book.douban.com/review/1029715/">《霍乱时期的爱情》——关于爱情的百科全书</a></h2>

                <div id="review_1029715_short" class="review-short" data-rid="1029715">
                    <div class="short-content">
                            <p class="spoiler-tip">这篇书评可能有关键情节透露</p>

                        对这部小说，我事先完全没有提防，甚至看到题目如同畅销小说般名叫《霍乱时期的爱情》的时候，我也没有放在心上，但到看完的时候，我惊诧了，惊诧于两件事：一是马尔克斯居然真的跑去写爱情，二是马尔克斯居然可以一点都不魔幻现实主义的。在这篇“魔幻色彩降到0”的小说里，他...

                        &nbsp;(<a href="javascript:;" id="toggle-1029715-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_1029715_full" class="hidden">
                    <div id="review_1029715_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="1029715" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png">
                        <span id="r-useful_count-1029715">
                                173
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="1029715" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png">
                        <span id="r-useless_count-1029715">
                                4
                        </span>
                    </a>
                    <a href="https://book.douban.com/review/1029715/#comments" class="reply ">21回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>




    

    

    <script type="text/javascript" src="https://img3.doubanio.com/misc/mixed_static/3f3c0900684b5500.js"></script>
    <!-- COLLECTED CSS -->
</div>








                <p class="pl">
                    &gt;
                    <a href="reviews">
                        更多书评
                            4005篇
                    </a>
                </p>
    </section>
<!-- COLLECTED JS -->

  









<div class="ugc-mod reading-notes">
  <div class="hd">
    <h2>
      读书笔记&nbsp;&nbsp;· · · · · ·&nbsp;
          <span class="pl">(<a href="https://book.douban.com/subject/10594787/annotation">共<span property="v:count">9447</span>篇</a>)</span>

    </h2>

        <a class="redbutt rr j a_show_login" href="https://www.douban.com/register?reason=annotate" rel="nofollow">
          <span>我来写笔记</span>
        </a>
    
  </div>
  

    <div class="bd">
      <ul class="inline-tabs">
        <li class="on"><a href="#rank" id="by_rank">按有用程度</a></li>
        <li><a href="#page" id="by_page">按页码先后</a></li>
        <li><a href="#time" id="by_time">最新笔记</a></li>
      </ul>
      
  <ul class="comments by_rank">
      
      <li class="ctsh clearfix" data-cid="20546089">
        <div class="ilst">
          <a href="https://www.douban.com/people/62644880/"><img src="https://img1.doubanio.com/icon/u62644880-17.jpg" alt="牧汀菊宝贝" class=""></a>
        </div>
        <div class="con">
          <div class="nlst">
            <h3>
              <div class="note-toggle rr">
                <a href="https://book.douban.com/annotation/20546089/" class="note-unfolder">展开</a>
                <a href="javascript:void(0);" class="note-folder">收起</a>
              </div>
              <a href="https://book.douban.com/annotation/20546089/" class="">《霍乱时期的爱情》最后一章读完后的感受</a></h3>
          </div>
          <div class="clst">
            <p class="user"><a href="https://www.douban.com/people/62644880/" class=" " title="牧汀菊宝贝">牧汀菊宝贝</a>
                (321，爱就像蓝天白云晴空万里)
              
                <span class="allstar50" title="力荐"></span>
            </p>
            <div class="reading-note" data-cid="20546089">
              <div class="short">
                
                  <span class="">如果视觉的镜头慢慢拉长，一眼望去穷尽一生，那种及生至死的缓和感就会令自己觉醒，原来在十岁经历的父母离异，或是二十岁时和初恋男友分手，更或是许许多多某一刻的痛彻心扉不堪回首，在这些事情里的感受全在时间的洪波里不足挂齿，总会被片片冲蚀掉，总会被带走，只有生活着的过程才是永恒。 这就是我读这本书最大的感受。 马尔克斯讲了两种爱情，一种因充满荷尔蒙悸动的一瞥而一眼万年，一种用家世背景相貌而量身打造却在漫...</span>
                  &nbsp;&nbsp;<a href="https://book.douban.com/annotation/20546089/">(<span class="">84回应</span>)</a>
                <p class="pl">
                  <span class="">2012-08-22 10:21</span>
                    &nbsp;&nbsp;<span class="">1591人喜欢</span>
                </p>
              </div>
              <div class="all hidden" style="display:none">
                <p>      如果视觉的镜头慢慢拉长，一眼望去穷尽一生，那种及生至死的缓和感就会令自己觉醒，原来在十岁经历的父母离异，或是二十岁时和初恋男友分手，更或是许许多多某一刻的痛彻心扉不堪回首，在这些事情里的感受全在时间的洪波里不足挂齿，总会被片片冲蚀掉，总会被带走，只有生活着的过程才是永恒。
   这就是我读这本书最大的感受。
   马尔克斯讲了两种爱情，一种因充满荷尔蒙悸动的一瞥而一眼万年，一种用家世背景相貌而量身打造却在漫长共度的时光里而历久弥新焦不离孟。哪一种才是生命永恒的主题。讲故事的人也说不清楚。所有美好的爱与痛苦都令人难以割舍，所以故事里的女人才先选择了生活中的爱情，又在丈夫去世后，登上理想中的爱情的轮船。
   这才发现，讨论哪种爱情是生命永恒主题没有意义，因为生命永恒的主题只是生命本身。只有生活着，生活的过程带来永恒的幸福，要告别人世的时候，才觉得自己的生命和心一直丰满又丰盛。</p>
                  <div class="col-rec-con clearfix">
                    







<div class="rec-sec">

    <span class="rec">

<a href="https://www.douban.com/accounts/register?reason=collect" class="j a_show_login lnk-sharing lnk-douban-sharing">推荐</a>
</span>
</div>

                  </div>
                <div class="pl col-time">
                  <a href="https://book.douban.com/annotation/20546089/#comments" class="">84回应</a>&nbsp;&nbsp;
                  2012-08-22 10:21
                </div>
              </div>
            </div>
          </div>
        </div>
      </li>
      
      <li class="ctsh clearfix" data-cid="21176754">
        <div class="ilst">
          <a href="https://www.douban.com/people/56641708/"><img src="https://img2.doubanio.com/icon/u56641708-2.jpg" alt="慢歌3" class=""></a>
        </div>
        <div class="con">
          <div class="nlst">
            <h3>
              <div class="note-toggle rr">
                <a href="https://book.douban.com/annotation/21176754/" class="note-unfolder">展开</a>
                <a href="javascript:void(0);" class="note-folder">收起</a>
              </div>
              <a href="https://book.douban.com/annotation/21176754/" class="">第379页</a></h3>
          </div>
          <div class="clst">
            <p class="user"><a href="https://www.douban.com/people/56641708/" class=" " title="慢歌3">慢歌3</a>
              
            </p>
            <div class="reading-note" data-cid="21176754">
              <div class="short">
                
                  <span class="">当一个女人决定和一个男人睡觉时，就没有她跃不过去的围墙，没有她推不倒的堡垒，也没有她抛不下的道德顾虑，事实上没有能管得住她的上帝。 谨慎做这样的决定。</span>
                  &nbsp;&nbsp;<a href="https://book.douban.com/annotation/21176754/">(<span class="">84回应</span>)</a>
                <p class="pl">
                  <span class="">2012-09-18 14:22</span>
                    &nbsp;&nbsp;<span class="">696人喜欢</span>
                </p>
              </div>
              <div class="all hidden" style="display:none">
                <figure>当一个女人决定和一个男人睡觉时，就没有她跃不过去的围墙，没有她推不倒的堡垒，也没有她抛不下的道德顾虑，事实上没有能管得住她的上帝。<figcaption>引自第379页</figcaption></figure><p>谨慎做这样的决定。</p>
                  <div class="col-rec-con clearfix">
                    







<div class="rec-sec">

    <span class="rec">

<a href="https://www.douban.com/accounts/register?reason=collect" class="j a_show_login lnk-sharing lnk-douban-sharing">推荐</a>
</span>
</div>

                  </div>
                <div class="pl col-time">
                  <a href="https://book.douban.com/annotation/21176754/#comments" class="">84回应</a>&nbsp;&nbsp;
                  2012-09-18 14:22
                </div>
              </div>
            </div>
          </div>
        </div>
      </li>
      
      <li class="ctsh clearfix" data-cid="25470699">
        <div class="ilst">
          <a href="https://www.douban.com/people/67393581/"><img src="https://img2.doubanio.com/icon/u67393581-3.jpg" alt="梦里花落知多少" class=""></a>
        </div>
        <div class="con">
          <div class="nlst">
            <h3>
              <div class="note-toggle rr">
                <a href="https://book.douban.com/annotation/25470699/" class="note-unfolder">展开</a>
                <a href="javascript:void(0);" class="note-folder">收起</a>
              </div>
              <a href="https://book.douban.com/annotation/25470699/" class="">第1页</a></h3>
          </div>
          <div class="clst">
            <p class="user"><a href="https://www.douban.com/people/67393581/" class=" " title="梦里花落知多少">梦里花落知多少</a>
              
            </p>
            <div class="reading-note" data-cid="25470699">
              <div class="short">
                
                  <span class="">真正的爱情需要什么？需要两个人在一起是轻松快乐的，没有压力。 . 爱一个人就是毫无保留地付出吗？ 不是。 每一个人都是一个独立的人，我们首先是属于自己的， 我们有思想，我们有个性，而不是把我们的全部都给对方。 我们可以有保留，比如你不愿意说的隐私，有秘密的人才是成熟的，不是吗？ 有时候不说出来反而更好。 . 真正成熟的选择，会选择和自己性格合适的人，能和自己一起过日子的人。 . 喜欢一个人，太急切了，反而不...</span>
                  &nbsp;&nbsp;<a href="https://book.douban.com/annotation/25470699/">(<span class="">24回应</span>)</a>
                <p class="pl">
                  <span class="">2013-03-30 01:38</span>
                    &nbsp;&nbsp;<span class="">541人喜欢</span>
                </p>
              </div>
              <div class="all hidden" style="display:none">
                <p>真正的爱情需要什么？需要两个人在一起是轻松快乐的，没有压力。
　　.
　　爱一个人就是毫无保留地付出吗？
　　不是。
　　每一个人都是一个独立的人，我们首先是属于自己的，
　　我们有思想，我们有个性，而不是把我们的全部都给对方。
　　我们可以有保留，比如你不愿意说的隐私，有秘密的人才是成熟的，不是吗？
　　有时候不说出来反而更好。
　　.
　 真正成熟的选择，会选择和自己性格合适的人，能和自己一起过日子的人。
　　.
　　喜欢一个人，太急切了，反而不好。
　　一是因为越想得到的越得不到；
　　二是得到了也很难珍惜，来得快去得也快。
　　细水长流一些，爱情会更长久。
　　.
　　相爱容易相处难。
　　相处中最重要的是宽容和妥协，在信任和了解的基础上。
　　没有宽容和妥协，任何两个人都无法相处。
　　。
　　纯纯的爱也许只有一次，但是真爱未必只有一次。
　　时间会抚平一切伤痕。
　　。
　　我们其实是可以爱上很多人的。
　　我们不是喜欢某个人，而是喜欢某种类型的人。
　　先来的人和我们相遇了，于是我们幸福地走到了一起；
　　对于后到的人，只能抱以歉意，同时，祝福他早日找到属于他自己的幸福。
　　没有谁是我们一生非拥有不可的，
　　爱一个人，很多时候实际上是习惯了这个人 。
　　。
　　现实和浪漫哪个更重要？
　　现实。
　　没有现实为基础，浪漫就是空中楼阁。
　　大学校园的爱情往往随着毕业而告终，大多是因为不现实，不在一个城市。
　　只有相互欣赏相互佩服各有所长的人，
　　才会碰撞出最美丽的火花，也才会结出最甜美的爱情果实。
　　爱情不等于生活，只是生活的一部分。
　　
　　。
　　恋爱的时间能长尽量长。
　　这最少有两点好处：
　　一，充分、尽可能长的享受恋爱的愉悦，
　　二，两人相处时间越长，越能检验彼此是否真心，越能看出两人性格是否合得来。
　　。
　　想知道一个人爱不爱你，就看他和你在一起有没有活力，
　　开不开心，有就是爱，没有就是不爱。
　　爱情不是感动，
　　你不是他心目中的理想伴侣，即使一时接受你，
　　将来碰上他心仪的那一位，一样会离开你。
　　。
　　浪漫是什么？
　　是送花？雨中漫步？楼前伫立不去？
　　如果两人彼此倾心相爱，什么事都不做，静静相对都会感觉是浪漫的。
　　否则，即使两人坐到月亮上拍拖，也是感觉不到浪漫的。
　　。
　　是否门当户对不要紧，最重要应该是兴当趣对，不然没有共同语言，
　　即使在一起，仍然会感觉到孤独。
　　。
　　持久的爱情源于彼此发自内心的真爱，建立在平等的基础之上。
　　任何只顾疯狂爱人而不顾自己有否被爱，
　　或是只顾享受被爱而不知真心爱人的人都不会有好的结局。
　　。
　　爱情既是风险投资，难免有去无回，失恋是再正常不过的事情。
　　爱过，就够了。
　　既然不能在一起，总有不能在一起的理由。
　　不能因为别人负了你，就不负责任地游戏、报复或是堕落，
　　自己演的戏，总要自己收场的。
　　何况，他不爱你，你做什么他都不会在乎。
　　。
　　如果爱上，就不要轻易放过机会。
　　莽撞，可能使你后悔一阵子；
　　怯懦，却可能使你后悔一辈子。
　　。
　　没有经历过爱情的人生是不完整的，没有经历过痛苦的爱情是不深刻的。
　　爱情使人生丰富，痛苦使爱情升华。
　　。
　　你可能习惯与现在的恋人，
　　明明不太喜欢，但在一起久了，习惯使人不太愿做新的选择。
　　人生会面临无数次选择。
　　当给你机会选择时，你一定要谨慎；
　　一旦你做出了选择，就永远不要后悔；
　　拿得起，放得下，该断则断，该忘记的，就把它忘记；
　　该珍惜的，就要把它珍惜
　　。
　　我们总说：“我要找一个很爱很爱的人。”
　　但是当对方问你，怎样才算是很爱很爱的时候，你却无法回答他，
　　因为你自己也不知道。
　　其实，很爱很爱的感觉，是要在一起经历了许多事情之后才会发现的。
　　　　。
　　当你爱一个人的时候，爱到八分绝对刚刚好。
　　所有的期待和希望都只有七八分，剩下两三分用来爱自己。
　　如果你还继续爱得更多，很可能会给对方沉重的压力，让彼此喘不过气来，
　　完全丧失了爱情的乐趣。
　　所以请记住，喝酒不要超过六分醉，吃饭不要超过七分饱，爱一个人不要超过八分。
　　。
　　爱一个人，要了解也要开解；
　　要道歉也要道谢；要认错也要改错；
　　要体贴也要体谅；是接受而不是忍受；
　　是宽容而不是纵容；是支持而不是支配；
　　是慰问而不是质问；是倾诉而不是控诉；
　　是难忘而不是遗忘；是彼此交流而不是凡事交代；
　　是为对方默默祈求而不是向对方诸多要求。
　　可以浪漫，但不要浪费，不要随便牵手，更不要随便放手。
　　。
　　浪漫的人这样描述与爱人的相逢：
　　千万人当中，在时间的无涯的荒野里，没有早一步，也没 有晚一步，刚巧赶上了。
　　两个人好着的时候，你不妨就这样想吧。
　　如果不好了，你要明白是否和某人在一起，不过是一个再简单不过的概率问题。
　　数千个擦肩而过中，你给谁机会谁就和你有缘分，纵没有甲，也会有乙。
　　别傻等那种想像中的木石前盟般的缘分了，生活中哪有那么多传奇。
　　别醒着做梦了，难道你忘了艺术虽然来源生活，却还高于生活吗？ 如果一只蚂蚁你对它好，它也能感觉得到。 你心里的善会让你形成一个令人舒服的气场，会形成一个人的气质。这是骗得了自己也骗不了别人的东西。所以，佛说，万事皆空善不空。无论什么时候，对人，对事，对畜生，对才狼虎豹都要心存善念。 人心就是人的命运，命运掌握在自己手中。</p>
                  <div class="col-rec-con clearfix">
                    







<div class="rec-sec">

    <span class="rec">

<a href="https://www.douban.com/accounts/register?reason=collect" class="j a_show_login lnk-sharing lnk-douban-sharing">推荐</a>
</span>
</div>

                  </div>
                <div class="pl col-time">
                  <a href="https://book.douban.com/annotation/25470699/#comments" class="">24回应</a>&nbsp;&nbsp;
                  2013-03-30 01:38
                </div>
              </div>
            </div>
          </div>
        </div>
      </li>
      
      <li class="ctsh clearfix" data-cid="20963019">
        <div class="ilst">
          <a href="https://www.douban.com/people/liunianshui/"><img src="https://img3.doubanio.com/icon/u3055239-10.jpg" alt="白洛之" class=""></a>
        </div>
        <div class="con">
          <div class="nlst">
            <h3>
              <div class="note-toggle rr">
                <a href="https://book.douban.com/annotation/20963019/" class="note-unfolder">展开</a>
                <a href="javascript:void(0);" class="note-folder">收起</a>
              </div>
              <a href="https://book.douban.com/annotation/20963019/" class="">第231页</a></h3>
          </div>
          <div class="clst">
            <p class="user"><a href="https://www.douban.com/people/liunianshui/" class=" " title="白洛之">白洛之</a>
                (垚之高，鋆熹好。)
              
                <span class="allstar40" title="推荐"></span>
            </p>
            <div class="reading-note" data-cid="20963019">
              <div class="short">
                
                  <span class="">这世界上到处都是幸福的寡妇。他曾看见她们在丈夫的尸体前痛苦得发疯，恳求别人把自己也放入同一口棺木，活活埋入地下，以免独自面对前路无法预知的苦难。可随着她们接受了现实，适应了新的境况，人们就会看到她们从尘土中站起来，获得新生。起初她们像阴影中的寄生虫一样生活在空荡荡的大房子里，向女仆们倾诉着心声，整日赖在枕头上：当了那么多年无所事事的囚徒，她们不知道自己该干些什么。为了打发绰绰有余的时间，她们为...</span>
                  &nbsp;&nbsp;<a href="https://book.douban.com/annotation/20963019/">(<span class="">17回应</span>)</a>
                <p class="pl">
                  <span class="">2012-09-08 23:55</span>
                    &nbsp;&nbsp;<span class="">390人喜欢</span>
                </p>
              </div>
              <div class="all hidden" style="display:none">
                <p>这世界上到处都是幸福的寡妇。他曾看见她们在丈夫的尸体前痛苦得发疯，恳求别人把自己也放入同一口棺木，活活埋入地下，以免独自面对前路无法预知的苦难。可随着她们接受了现实，适应了新的境况，人们就会看到她们从尘土中站起来，获得新生。起初她们像阴影中的寄生虫一样生活在空荡荡的大房子里，向女仆们倾诉着心声，整日赖在枕头上：当了那么多年无所事事的囚徒，她们不知道自己该干些什么。为了打发绰绰有余的时间，她们为死者的衣服钉上以前从来没有时间去钉的扣子，把他们的衬衫熨了又熨，还给袖口和领口上蜡，让它们时刻保持完美。她们继续为死去的丈夫在浴室放上香皂，在床上铺好带有他们名字首字母的床罩，在餐桌他们的位置上摆好餐具，以防死者说不定什么时候没有事先通知就回来了，就像他们生前常做的那样。但当她们独自去望弥撒时，才逐渐意识到，自己又一次成为自己意愿的主人，当初，为了换取一种安全感，她们不仅放弃了自己家庭的姓氏，甚至放弃了自我，可那种安全感不过是她们做姑娘时许多幻想中的一个罢了。只有她们自己知道，她们曾经疯狂爱着的那个男人——尽管他或许也爱着她们——给她们带来的负担有多么沉重，她们不得不照顾他们直到最后一口气，喂他们吃喝，给他们换下脏兮兮的尿布，用母亲式的巧妙花招哄他们开心，以减轻他们清晨走出家门去直面现实的恐惧。可当看到他们受自己的鼓动离开家门，准备一口去吞掉整个世界时，她们又开始害怕男人会一去不复返。这就是生活。而爱，如果真的存在，则是另一回事：另一种生活。
然而，在孤独中休养生息时，寡妇们发现，诚实的生活方式其实是按照自己身体的意愿行事，饿的时候才吃饭，爱的时候不必撒谎，睡觉的时候也不用为了逃避可耻的爱情程式而装睡，自己终于成了整张床的主人，它的全部都归自己独享，再没有人跟她们争一半的床单、一半的空气和一半的夜晚，甚至身体也终于能尽情做属于自己的梦，能自然而然地独自醒来了。</p>
                  <div class="col-rec-con clearfix">
                    







<div class="rec-sec">

    <span class="rec">

<a href="https://www.douban.com/accounts/register?reason=collect" class="j a_show_login lnk-sharing lnk-douban-sharing">推荐</a>
</span>
</div>

                  </div>
                <div class="pl col-time">
                  <a href="https://book.douban.com/annotation/20963019/#comments" class="">17回应</a>&nbsp;&nbsp;
                  2012-09-08 23:55
                </div>
              </div>
            </div>
          </div>
        </div>
      </li>
  </ul>
  

      
  <ul class="comments by_page" style="display: none">
      
      <li class="ctsh clearfix" data-cid="30106307">
        <div class="ilst">
          <a href="https://www.douban.com/people/61714932/"><img src="https://img1.doubanio.com/icon/user_normal.jpg" alt="三碟连放" class=""></a>
        </div>
        <div class="con">
          <div class="nlst">
            <h3>
              <div class="note-toggle rr">
                <a href="https://book.douban.com/annotation/30106307/" class="note-unfolder">展开</a>
                <a href="javascript:void(0);" class="note-folder">收起</a>
              </div>
              <a href="https://book.douban.com/annotation/30106307/" class="">第-16277页</a></h3>
          </div>
          <div class="clst">
            <p class="user"><a href="https://www.douban.com/people/61714932/" class=" " title="三碟连放">三碟连放</a>
              
            </p>
            <div class="reading-note" data-cid="30106307">
              <div class="short">
                
                  <span class="">腐败读的</span>
                <p class="pl">
                  <span class="">2014-01-11 02:13</span>
                </p>
              </div>
              <div class="all hidden" style="display:none">
                <p>腐败读的</p>
                  <div class="col-rec-con clearfix">
                    







<div class="rec-sec">

    <span class="rec">

<a href="https://www.douban.com/accounts/register?reason=collect" class="j a_show_login lnk-sharing lnk-douban-sharing">推荐</a>
</span>
</div>

                  </div>
                <div class="pl col-time">
                  <a href="https://book.douban.com/annotation/30106307/#comments" class="">回应</a>&nbsp;&nbsp;
                  2014-01-11 02:13
                </div>
              </div>
            </div>
          </div>
        </div>
      </li>
      
      <li class="ctsh clearfix" data-cid="20832696">
        <div class="ilst">
          <a href="https://www.douban.com/people/54897400/"><img src="https://img1.doubanio.com/icon/user_normal.jpg" alt="lier" class=""></a>
        </div>
        <div class="con">
          <div class="nlst">
            <h3>
              <div class="note-toggle rr">
                <a href="https://book.douban.com/annotation/20832696/" class="note-unfolder">展开</a>
                <a href="javascript:void(0);" class="note-folder">收起</a>
              </div>
              <a href="https://book.douban.com/annotation/20832696/" class="">第1页</a></h3>
          </div>
          <div class="clst">
            <p class="user"><a href="https://www.douban.com/people/54897400/" class=" " title="lier">lier</a>
              
            </p>
            <div class="reading-note" data-cid="20832696">
              <div class="short">
                
                  <span class="">不可避免，苦杏仁的气味总是让他想起爱情受阻的命运。</span>
                <p class="pl">
                  <span class="">2012-09-03 16:29</span>
                    &nbsp;&nbsp;<span class="">3人喜欢</span>
                </p>
              </div>
              <div class="all hidden" style="display:none">
                <p>不可避免，苦杏仁的气味总是让他想起爱情受阻的命运。</p>
                  <div class="col-rec-con clearfix">
                    







<div class="rec-sec">

    <span class="rec">

<a href="https://www.douban.com/accounts/register?reason=collect" class="j a_show_login lnk-sharing lnk-douban-sharing">推荐</a>
</span>
</div>

                  </div>
                <div class="pl col-time">
                  <a href="https://book.douban.com/annotation/20832696/#comments" class="">回应</a>&nbsp;&nbsp;
                  2012-09-03 16:29
                </div>
              </div>
            </div>
          </div>
        </div>
      </li>
      
      <li class="ctsh clearfix" data-cid="20889176">
        <div class="ilst">
          <a href="https://www.douban.com/people/1625628/"><img src="https://img3.doubanio.com/icon/u1625628-30.jpg" alt="恰零几" class=""></a>
        </div>
        <div class="con">
          <div class="nlst">
            <h3>
              <div class="note-toggle rr">
                <a href="https://book.douban.com/annotation/20889176/" class="note-unfolder">展开</a>
                <a href="javascript:void(0);" class="note-folder">收起</a>
              </div>
              <a href="https://book.douban.com/annotation/20889176/" class="">第1页</a></h3>
          </div>
          <div class="clst">
            <p class="user"><a href="https://www.douban.com/people/1625628/" class=" " title="恰零几">恰零几</a>
                (啊，浅薄的部落，怎能不可怜你们)
              
                <span class="allstar50" title="力荐"></span>
            </p>
            <div class="reading-note" data-cid="20889176">
              <div class="short">
                
                  <span class="">回头写</span>
                <p class="pl">
                  <span class="">2012-09-05 21:28</span>
                </p>
              </div>
              <div class="all hidden" style="display:none">
                <p>回头写
</p>
                  <div class="col-rec-con clearfix">
                    







<div class="rec-sec">

    <span class="rec">

<a href="https://www.douban.com/accounts/register?reason=collect" class="j a_show_login lnk-sharing lnk-douban-sharing">推荐</a>
</span>
</div>

                  </div>
                <div class="pl col-time">
                  <a href="https://book.douban.com/annotation/20889176/#comments" class="">回应</a>&nbsp;&nbsp;
                  2012-09-05 21:28
                </div>
              </div>
            </div>
          </div>
        </div>
      </li>
      
      <li class="ctsh clearfix" data-cid="21032824">
        <div class="ilst">
          <a href="https://www.douban.com/people/53262550/"><img src="https://img1.doubanio.com/icon/user_normal.jpg" alt="蠹鱼" class=""></a>
        </div>
        <div class="con">
          <div class="nlst">
            <h3>
              <div class="note-toggle rr">
                <a href="https://book.douban.com/annotation/21032824/" class="note-unfolder">展开</a>
                <a href="javascript:void(0);" class="note-folder">收起</a>
              </div>
              <a href="https://book.douban.com/annotation/21032824/" class="">第1页</a></h3>
          </div>
          <div class="clst">
            <p class="user"><a href="https://www.douban.com/people/53262550/" class=" " title="蠹鱼">蠹鱼</a>
                (我们读书吧)
              
            </p>
            <div class="reading-note" data-cid="21032824">
              <div class="short">
                
                  <span class="">放下所有的书，来看这本。</span>
                <p class="pl">
                  <span class="">2012-09-11 22:44</span>
                </p>
              </div>
              <div class="all hidden" style="display:none">
                <p>放下所有的书，来看这本。</p>
                  <div class="col-rec-con clearfix">
                    







<div class="rec-sec">

    <span class="rec">

<a href="https://www.douban.com/accounts/register?reason=collect" class="j a_show_login lnk-sharing lnk-douban-sharing">推荐</a>
</span>
</div>

                  </div>
                <div class="pl col-time">
                  <a href="https://book.douban.com/annotation/21032824/#comments" class="">回应</a>&nbsp;&nbsp;
                  2012-09-11 22:44
                </div>
              </div>
            </div>
          </div>
        </div>
      </li>
  </ul>
  

      
  <ul class="comments by_time" style="display: none">
      
      <li class="ctsh clearfix" data-cid="104427307">
        <div class="ilst">
          <a href="https://www.douban.com/people/176008797/"><img src="https://img3.doubanio.com/icon/u176008797-1.jpg" alt="之一" class=""></a>
        </div>
        <div class="con">
          <div class="nlst">
            <h3>
              <div class="note-toggle rr">
                <a href="https://book.douban.com/annotation/104427307/" class="note-unfolder">展开</a>
                <a href="javascript:void(0);" class="note-folder">收起</a>
              </div>
              <a href="https://book.douban.com/annotation/104427307/" class="">世俗的真相</a></h3>
          </div>
          <div class="clst">
            <p class="user"><a href="https://www.douban.com/people/176008797/" class=" " title="之一">之一</a>
              
            </p>
            <div class="reading-note" data-cid="104427307">
              <div class="short">
                
                  <span class="">社交生活的关键在于学会控制恐惧，夫妻生活的关键在于学会控制厌恶。</span>
                <p class="pl">
                  <span class="">2021-03-20 23:24</span>
                </p>
              </div>
              <div class="all hidden" style="display:none">
                <p>社交生活的关键在于学会控制恐惧，夫妻生活的关键在于学会控制厌恶。</p>
                  <div class="col-rec-con clearfix">
                    







<div class="rec-sec">

    <span class="rec">

<a href="https://www.douban.com/accounts/register?reason=collect" class="j a_show_login lnk-sharing lnk-douban-sharing">推荐</a>
</span>
</div>

                  </div>
                <div class="pl col-time">
                  <a href="https://book.douban.com/annotation/104427307/#comments" class="">回应</a>&nbsp;&nbsp;
                  2021-03-20 23:24
                </div>
              </div>
            </div>
          </div>
        </div>
      </li>
      
      <li class="ctsh clearfix" data-cid="104416693">
        <div class="ilst">
          <a href="https://www.douban.com/people/206822567/"><img src="https://img9.doubanio.com/icon/u206822567-26.jpg" alt="栗子啊哩zi" class=""></a>
        </div>
        <div class="con">
          <div class="nlst">
            <h3>
              <div class="note-toggle rr">
                <a href="https://book.douban.com/annotation/104416693/" class="note-unfolder">展开</a>
                <a href="javascript:void(0);" class="note-folder">收起</a>
              </div>
              <a href="https://book.douban.com/annotation/104416693/" class="">第48页</a></h3>
          </div>
          <div class="clst">
            <p class="user"><a href="https://www.douban.com/people/206822567/" class=" " title="栗子啊哩zi">栗子啊哩zi</a>
              
            </p>
            <div class="reading-note" data-cid="104416693">
              <div class="short">
                
                  <span class="">丈夫已经奄奄一息，但还在坚持与死神这致命的一击做着最后一分钟抗争，好让她及时赶来。要这样撇下她独自离去，他感到无比痛苦，透过泪水，他在慌乱的人群中认出了她。他诀别地看了她最后一眼，在两人半个世纪的共同生活中，她从未见过他的眼神如此闪亮，如此悲痛，而又如此充满感激。他用尽最后一口气，对她说道： “只有上帝知道我有多爱你。”</span>
                <p class="pl">
                  <span class="">2021-03-20 16:47</span>
                </p>
              </div>
              <div class="all hidden" style="display:none">
                <figure>丈夫已经奄奄一息，但还在坚持与死神这致命的一击做着最后一分钟抗争，好让她及时赶来。要这样撇下她独自离去，他感到无比痛苦，透过泪水，他在慌乱的人群中认出了她。他诀别地看了她最后一眼，在两人半个世纪的共同生活中，她从未见过他的眼神如此闪亮，如此悲痛，而又如此充满感激。他用尽最后一口气，对她说道：
    “只有上帝知道我有多爱你。”<figcaption>引自第48页</figcaption></figure>
                  <div class="col-rec-con clearfix">
                    







<div class="rec-sec">

    <span class="rec">

<a href="https://www.douban.com/accounts/register?reason=collect" class="j a_show_login lnk-sharing lnk-douban-sharing">推荐</a>
</span>
</div>

                  </div>
                <div class="pl col-time">
                  <a href="https://book.douban.com/annotation/104416693/#comments" class="">回应</a>&nbsp;&nbsp;
                  2021-03-20 16:47
                </div>
              </div>
            </div>
          </div>
        </div>
      </li>
      
      <li class="ctsh clearfix" data-cid="104416550">
        <div class="ilst">
          <a href="https://www.douban.com/people/83398513/"><img src="https://img2.doubanio.com/icon/u83398513-3.jpg" alt="几乎是幻想" class=""></a>
        </div>
        <div class="con">
          <div class="nlst">
            <h3>
              <div class="note-toggle rr">
                <a href="https://book.douban.com/annotation/104416550/" class="note-unfolder">展开</a>
                <a href="javascript:void(0);" class="note-folder">收起</a>
              </div>
              <a href="https://book.douban.com/annotation/104416550/" class="">第242页</a></h3>
          </div>
          <div class="clst">
            <p class="user"><a href="https://www.douban.com/people/83398513/" class=" " title="几乎是幻想">几乎是幻想</a>
              
            </p>
            <div class="reading-note" data-cid="104416550">
              <div class="short">
                
                  <span class="">所谓的世俗生活，虽然在她了解之前曾有过许多疑虑，但其实那不过是一套沿自传统的规矩，庸俗的仪式，事先想好的言词，在此之下，人们彼此消遣，为的是不致互相杀戮。在这个轻浮的世俗天堂，最显著的特征就是对陌生事物的恐惧。她用一种更为简单的方式为它下了定义：“社交生活的关键在于学会控制恐惧，夫妻生活的关键在于学会控制厌恶。”</span>
                <p class="pl">
                  <span class="">2021-03-20 16:41</span>
                </p>
              </div>
              <div class="all hidden" style="display:none">
                <figure>所谓的世俗生活，虽然在她了解之前曾有过许多疑虑，但其实那不过是一套沿自传统的规矩，庸俗的仪式，事先想好的言词，在此之下，人们彼此消遣，为的是不致互相杀戮。在这个轻浮的世俗天堂，最显著的特征就是对陌生事物的恐惧。她用一种更为简单的方式为它下了定义：“社交生活的关键在于学会控制恐惧，夫妻生活的关键在于学会控制厌恶。”<figcaption>引自第242页</figcaption></figure>
                  <div class="col-rec-con clearfix">
                    







<div class="rec-sec">

    <span class="rec">

<a href="https://www.douban.com/accounts/register?reason=collect" class="j a_show_login lnk-sharing lnk-douban-sharing">推荐</a>
</span>
</div>

                  </div>
                <div class="pl col-time">
                  <a href="https://book.douban.com/annotation/104416550/#comments" class="">回应</a>&nbsp;&nbsp;
                  2021-03-20 16:41
                </div>
              </div>
            </div>
          </div>
        </div>
      </li>
      
      <li class="ctsh clearfix" data-cid="104416472">
        <div class="ilst">
          <a href="https://www.douban.com/people/206822567/"><img src="https://img9.doubanio.com/icon/u206822567-26.jpg" alt="栗子啊哩zi" class=""></a>
        </div>
        <div class="con">
          <div class="nlst">
            <h3>
              <div class="note-toggle rr">
                <a href="https://book.douban.com/annotation/104416472/" class="note-unfolder">展开</a>
                <a href="javascript:void(0);" class="note-folder">收起</a>
              </div>
              <a href="https://book.douban.com/annotation/104416472/" class="">第46页</a></h3>
          </div>
          <div class="clst">
            <p class="user"><a href="https://www.douban.com/people/206822567/" class=" " title="栗子啊哩zi">栗子啊哩zi</a>
              
            </p>
            <div class="reading-note" data-cid="104416472">
              <div class="short">
                
                  <span class="">八十一岁时，他仍旧足够清醒地意识到，把自已拴在这个世界上的，仅剩下几根细细的丝线，睡梦中简单地改变一下姿势都可能让它们毫无痛苦地断开。而如果说，他还在尽可能地维持它们，那完全是出于在死亡的黑暗中找不到上帝的恐惧。</span>
                <p class="pl">
                  <span class="">2021-03-20 16:39</span>
                </p>
              </div>
              <div class="all hidden" style="display:none">
                <figure>八十一岁时，他仍旧足够清醒地意识到，把自已拴在这个世界上的，仅剩下几根细细的丝线，睡梦中简单地改变一下姿势都可能让它们毫无痛苦地断开。而如果说，他还在尽可能地维持它们，那完全是出于在死亡的黑暗中找不到上帝的恐惧。<figcaption>引自第46页</figcaption></figure>
                  <div class="col-rec-con clearfix">
                    







<div class="rec-sec">

    <span class="rec">

<a href="https://www.douban.com/accounts/register?reason=collect" class="j a_show_login lnk-sharing lnk-douban-sharing">推荐</a>
</span>
</div>

                  </div>
                <div class="pl col-time">
                  <a href="https://book.douban.com/annotation/104416472/#comments" class="">回应</a>&nbsp;&nbsp;
                  2021-03-20 16:39
                </div>
              </div>
            </div>
          </div>
        </div>
      </li>
  </ul>
  

    </div>
      <div class="ft">
        <p class="trr">&gt; <a href="https://book.douban.com/subject/10594787/annotation">更多读书笔记（共9447篇）</a></p>
      </div>

</div>



<script type="text/javascript">
  $(document).ready(function(){
    var TEMPL_ADD_COL = '<a href="" id="n-{NOTE_ID}" class="colbutt ll add-col"><span>收藏</span></a>',
      TEMPL_DEL_COL = '<span class="pl">已收藏 &gt;<a href="" id="n-{NOTE_ID}" class="del-col">取消收藏</a></span>';

    $('body').delegate('.add-col', 'click', function(e){
      e.preventDefault();
      var nnid = $(this).attr('id').split('-')[1],
        bn_add = $(this);
      $.post_withck("/j/book/annotation_collect",{nid:nnid},function(){
        var a = $(TEMPL_DEL_COL.replace('{NOTE_ID}', nnid));
        bn_add.before(a);
        bn_add.remove();
      })
    });

    $('body').delegate('.del-col', 'click', function(e){
      e.preventDefault();
      var nnid = $(this).attr('id').split('-')[1],
        bn_del = $(this).parent();
      $.post_withck("/j/book/annotation_uncollect", {nid: nnid}, function() {
        var a = $(TEMPL_ADD_COL.replace('{NOTE_ID}', nnid));
        bn_del.before(a);
        bn_del.remove();
      })
    });

    $("pre.source").each(function(){
      var cn = $(this).attr('class').split(' ');
      l = cn[1];
      s = 'rand01';
      n = cn[2];
      $(this).snippet(n,{style: s, showNum: l});
    });

    var annotationMod = $('.reading-notes .bd')
      , annotationTabs = annotationMod.find('.inline-tabs li')
      , annotationTabLinks = annotationTabs.find('a')
      , annotationTabContents = annotationMod.find('ul.comments');

    annotationTabLinks.click(function(e){
      e.preventDefault();
      var el = $(this)
        , kind = el.attr('id');

      annotationTabs.removeClass('on');
      el.parent().addClass('on');
      annotationTabContents.hide();
      annotationTabContents.filter('.' + kind).show();
    });
  });
</script>

<script type="text/x-mathjax-config">
MathJax.Hub.Config({
	jax: ["input/TeX", "output/HTML-CSS"],
    extensions: ["tex2jax.js","TeX/AMSmath.js","TeX/AMSsymbols.js","TeX/noUndefined.js"],
    tex2jax: {
		inlineMath: [ ["($", "$)"], ['\\(','\\)'] ],
		displayMath: [ ["($$","$$)"], ['\\[','\\]']],
		skipTags: ["script","noscript","style","textarea"],
		processEscapes: true,
		processEnvironments: true,
		preview: "TeX"
	},
	showProcessingMessages: false
  });
</script>


  






<div id="db-discussion-section" class="indent ugc-mod">




        
  

  <h2>
    <span>论坛</span>
      &nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·

  </h2>



        
<table class="olt"><tbody><tr><td></td><td></td><td></td><td></td></tr>
        
        <tr>
            <td class="pl"><a href="https://book.douban.com/subject/10594787/discussion/616898985/" title="霍乱时期的爱情主要人物及关系结构图">霍乱时期的爱情主要人物及关系结构图</a></td>
            <td class="pl"><span>来自</span><a href="https://www.douban.com/people/168260796/">padded</a></td>
            <td class="pl"><span>2 回应</span></td>
            <td class="pl"><span>2021-03-21</span></td>
        </tr>
        
        <tr>
            <td class="pl"><a href="https://book.douban.com/subject/10594787/discussion/616694469/" title="女主对她女儿有点太过分了">女主对她女儿有点太过分了</a></td>
            <td class="pl"><span>来自</span><a href="https://www.douban.com/people/170476627/">猫熊迷丶</a></td>
            <td class="pl"><span>3 回应</span></td>
            <td class="pl"><span>2021-03-21</span></td>
        </tr>
        
        <tr>
            <td class="pl"><a href="https://book.douban.com/subject/10594787/discussion/616747934/" title="大家都是在什么年龄读的这本书？">大家都是在什么年龄读的这本书？</a></td>
            <td class="pl"><span>来自</span><a href="https://www.douban.com/people/170184217/">SalmaBennani1</a></td>
            <td class="pl"><span>38 回应</span></td>
            <td class="pl"><span>2021-03-20</span></td>
        </tr>
        
        <tr>
            <td class="pl"><a href="https://book.douban.com/subject/10594787/discussion/615110140/" title="没人觉得和14岁女孩那段很猥琐吗？">没人觉得和14岁女孩那段很猥琐吗？</a></td>
            <td class="pl"><span>来自</span><a href="https://www.douban.com/people/48081014/">茱莉娅·萝卜丝</a></td>
            <td class="pl"><span>144 回应</span></td>
            <td class="pl"><span>2021-03-18</span></td>
        </tr>
        
        <tr>
            <td class="pl"><a href="https://book.douban.com/subject/10594787/discussion/616756724/" title="为什么达萨在市场遇到阿里萨以后突然就决定结束关系了？">为什么达萨在市场遇到阿里萨以后突然就决定结束关...</a></td>
            <td class="pl"><span>来自</span><a href="https://www.douban.com/people/217401456/">豆友217401456</a></td>
            <td class="pl"><span>21 回应</span></td>
            <td class="pl"><span>2021-03-15</span></td>
        </tr>
</tbody></table>


            <p class="pl" align="right">&gt;
                <a href="https://book.douban.com/subject/10594787/discussion/">浏览更多话题</a>
            </p>


</div>




</div>
<script type="text/javascript">
$(function(){if($.browser.msie && $.browser.version == 6.0){
    var maxWidth = parseInt($('#info').css('max-width'));
    if($('#info').width() > maxWidth)
        $('#info').width(maxWidth)
}});
</script>
</div>
      <div class="aside" style="position: static;">
        
  
  






  <div id="dale_book_subject_top_right" ad-status="appended" data-sell-type="RTB" data-type="GoogleRender" data-version="3.2.28"><div style="position: relative; margin: 0px 0px 20px; display: block; overflow: hidden;"><div id="dale_book_subject_top_right_frame"><script src="//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"></script><ins class="adsbygoogle" style="display:inline-block;width:300px;height:250px" data-ad-client="ca-pub-4830389020085397" data-ad-slot="1983604743" data-adsbygoogle-status="done"><ins id="aswift_1_expand" style="display:inline-table;border:none;height:250px;margin:0;padding:0;position:relative;visibility:visible;width:300px;background-color:transparent;" tabindex="0" title="Advertisement" aria-label="Advertisement"><ins id="aswift_1_anchor" style="display:block;border:none;height:250px;margin:0;padding:0;position:relative;visibility:visible;width:300px;background-color:transparent;"><iframe id="aswift_1" name="aswift_1" style="left:0;position:absolute;top:0;border:0;width:300px;height:250px;" sandbox="allow-forms allow-popups allow-popups-to-escape-sandbox allow-same-origin allow-scripts allow-top-navigation-by-user-activation" width="300" height="250" frameborder="0" src="https://googleads.g.doubleclick.net/pagead/ads?client=ca-pub-4830389020085397&amp;output=html&amp;h=250&amp;slotname=1983604743&amp;adk=4277889270&amp;adf=3367803850&amp;pi=t.ma~as.1983604743&amp;w=300&amp;lmt=1616331164&amp;psa=1&amp;format=300x250&amp;url=https%3A%2F%2Fbook.douban.com%2Fsubject%2F10594787%2F&amp;flash=0&amp;wgl=1&amp;uach=WyJtYWNPUyIsIjExXzJfMyIsIng4NiIsIiIsIjg5LjAuNDM4OS45MCIsW11d&amp;dt=1616331164034&amp;bpp=2&amp;bdt=835&amp;idt=2&amp;shv=r20210316&amp;cbv=r20190131&amp;ptt=9&amp;saldr=aa&amp;abxe=1&amp;cookie=ID%3D0907fc60228d7d09-22ebc0075ec50078%3AT%3D1609135239%3ART%3D1609135239%3AR%3AS%3DALNI_MYywACpt3YcT7bf33LHhYdUXS_xfQ&amp;prev_fmts=0x0&amp;nras=1&amp;correlator=2957303762013&amp;frm=20&amp;pv=2&amp;ga_vid=813602588.1616295532&amp;ga_sid=1616331164&amp;ga_hid=620866296&amp;ga_fc=1&amp;u_tz=480&amp;u_his=4&amp;u_java=0&amp;u_h=1050&amp;u_w=1680&amp;u_ah=1025&amp;u_aw=1636&amp;u_cd=30&amp;u_nplug=3&amp;u_nmime=4&amp;adx=712&amp;ady=254&amp;biw=559&amp;bih=914&amp;scr_x=0&amp;scr_y=0&amp;eid=21068946%2C44739387&amp;oid=3&amp;pvsid=2477251061253804&amp;pem=111&amp;ref=https%3A%2F%2Fbook.douban.com%2Ftag%2F%25E5%25B0%258F%25E8%25AF%25B4&amp;rx=0&amp;eae=0&amp;fc=1920&amp;brdim=-1629%2C25%2C-1629%2C25%2C1636%2C25%2C1629%2C1025%2C559%2C914&amp;vis=1&amp;rsz=%7C%7CoeE%7C&amp;abl=CS&amp;pfx=0&amp;fu=9216&amp;bc=31&amp;ifi=2&amp;uci=a!2&amp;fsb=1&amp;xpc=XRJKIRofOJ&amp;p=https%3A//book.douban.com&amp;dtd=11" marginwidth="0" marginheight="0" vspace="0" hspace="0" allowtransparency="true" scrolling="no" allowfullscreen="true" allow="conversion-measurement ‘src’" data-google-container-id="a!2" data-google-query-id="CNDjpr62we8CFUsdGAodkl8EZA" data-load-complete="true"></iframe></ins></ins></ins><script>
(adsbygoogle = window.adsbygoogle || []).push({});
</script><script src="https://erebor.douban.com/count/?ad=204134&amp;bid=SJ02sjYB7c8&amp;chicken=1cce5e06e158bc913fcd308ff3a6da74&amp;crtr=7%3A%E9%9C%8D%E4%B9%B1%E6%97%B6%E6%9C%9F%E7%9A%84%E7%88%B1%E6%83%85%7C7%3A%E7%88%B1%E6%83%85%7C7%3A%E5%8A%A0%E8%A5%BF%E4%BA%9A-%E9%A9%AC%E5%B0%94%E5%85%8B%E6%96%AF%7C7%3A%E5%8A%A0%E8%A5%BF%E4%BA%9A%C2%B7%E9%A9%AC%E5%B0%94%E5%85%8B%E6%96%AF%7C7%3A%E5%B0%8F%E8%AF%B4%7C7%3A%E6%8B%89%E7%BE%8E%E6%96%87%E5%AD%A6%7C7%3A%E5%A4%96%E5%9B%BD%E6%96%87%E5%AD%A6%7C7%3A%E6%96%87%E5%AD%A6%7C7%3A%E7%BB%8F%E5%85%B8%7C7%3A%E4%BA%BA%E7%94%9F%7C3%3A%2Fsubject%2F10594787%2F&amp;device=100&amp;experiment_id=0&amp;fv=&amp;hit_word=&amp;hn=fram44&amp;lat=0.00000&amp;lon=0.00000&amp;mark=&amp;model=&amp;net=&amp;ns=1616331073290179392&amp;os=51&amp;osv=X+11_2_3&amp;p=0&amp;pid=debug_5a495f621317c50507dbb577b671e1ebd20d5f39&amp;price=TITDjNM8_pF57taxshxwGg&amp;rexxar=0&amp;type=impression&amp;uid=&amp;unit=dale_book_subject_top_right&amp;user_variation=&amp;vendor="></script></div><div style="line-height: 1; text-align: center; background-color: rgba(0, 0, 0, 0.3); font-size: 12px; position: absolute; right: 0px; bottom: 0px; padding: 4px; color: rgb(255, 255, 255);">由谷歌提供的广告</div></div></div>
    
  
  
  <div class="gray_ad">
    
  

  <h2>
    <span>其他版本在豆瓣书店有售</span>
      &nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·

  </h2>


    <div id="impression_track_market_sale" data-track="https://frodo.douban.com/rohirrim/tracking/impression?subject_id=10594787&amp;isbn=7544258971&amp;source=douban_market&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject"></div>
    <div class="market-banner">
        <a class="publish" href="https://book.douban.com/subject/34924580/" target="_blank">
          南海出版公司版
        </a><br>
      <span class="title">
        纸质版&nbsp;
      </span>
      <span class="price"> 55.20元</span>
        <span class="price"> <s>69.00元</s></span>
      <span class="promotion-info">满48元包邮</span>
      <div class="actions">
          <a class="j a_show_login buy-btn buy" href="javascript:;">去购买</a>
          <a class="j a_show_login buy-btn cart" href="javascript:;">加入购物车</a>
      </div>
    </div>
  </div>
  <script>
    $(document).ready(function(){
      var marketSaleDOM = $('#impression_track_market_sale')
      if (marketSaleDOM) {
        var mstUrl = $('#impression_track_market_sale').attr('data-track')
        reportTrack(mstUrl)
      }

      $('#click_track_market_buy').click(function(e){
        var mbtUrl = e.target.dataset.track
        reportTrack(mbtUrl)
      })
      $('#click_track_market_cart').click(function(e){
        var mctUrl = e.target.dataset.track
        reportTrack(mctUrl)
      })
      function reportTrack(url) {
        if (!url) return false
        $.ajax({ url: url, dataType: 'text/html' })
      }
    })
  </script>


  






<style type="text/css" media="screen">
  .add2cartContainer{overflow:hidden;vertical-align:bottom;line-height:1;padding:10px 0}.add2cartContainer .add2cart{margin-right:0;display:inline-block}.buyinfo .bs{margin:0}.buyinfo .current-version-list{margin-bottom:27px}.buyinfo .secondhand-books-list{margin-top:-10px}.buyinfo .secondhand-books-list li{display:block}.buyinfo .secondhand-books-list .flex-wrap{display:flex}.buyinfo .secondhand-books-list .flex-wrap .vendor-name{flex:0 0 33%}.buyinfo .secondhand-books-list .flex-wrap .cell:last-child{width:70px;flex:initial;text-align:right}.buyinfo li{display:flex;flex-direction:column;padding:12px 0;position:relative;line-height:1;border-bottom:1px solid rgba(0,0,0,0.08)}.buyinfo li .cell{flex:1}.buyinfo li .cell:first-child a{line-height:20px}.buyinfo li .vendor-name.cell,.buyinfo li .vendor-name{flex:0 0 33%;align-self:normal;line-height:20px}.buyinfo li .price-btn-wrapper{display:flex}.buyinfo li .price-btn-wrapper>.cell{display:flex;margin-bottom:5px}.buyinfo li .price-btn-wrapper>.cell:last-child{margin-bottom:0}.buyinfo li .price-btn-wrapper>.cell .cell:last-child{padding-right:1px;width:70px;flex:initial;text-align:right}.buyinfo li .buylink-price{line-height:20px}.buyinfo li .buylink-price.promotion{color:#c34}.buyinfo li s{padding-left:2px;line-height:20px;font-family:Helvetica;font-size:11px;color:rgba(0,0,0,0.25);letter-spacing:0;vertical-align:top}.buyinfo li .publisher{color:rgba(0,0,0,0.25)}.buyinfo li .buy-book-btn{display:inline-block;height:20px;line-height:20px;width:70px;text-align:center;background:#fff;border-radius:3px;font-size:12px;box-shadow:0 0 1px 0 rgba(0,0,0,0.07)}.buyinfo li .buy-book-btn:hover{opacity:.8}.buyinfo li .secondhand-book-btn,.buyinfo li .paper-book-btn{color:#EB9108}.buyinfo li .secondhand-book-btn:hover,.buyinfo li .paper-book-btn:hover{background:#EB9108;color:#fff}.buyinfo li .e-book-btn{color:#32A998}.buyinfo li .e-book-btn:hover{background:#32A998;color:#fff}.buyinfo li .go-shopping{display:inline-block;line-height:20px;width:70px;text-align:center}.buyinfo .more-info{color:rgba(0,0,0,0.5);font-size:12px}.buyinfo .more-info.not-onsale{color:rgba(0,0,0,0.5)}.buyinfo .cell.not-onsale .more-info{line-height:20px}.buyinfo-printed{margin-bottom:15px}.buyinfo-printed.no-border{border-bottom:0}.buyinfo-printed .presale-indicator{margin:0;width:auto;text-indent:0;background:none;color:rgba(0,0,0,0.5)}.gray_ad{padding:18px 16px;background:#F6F6F2}.gray_ad h2{margin-bottom:10px;font-size:15px;color:#0B7C2A}.gray_ad .ebook-tag{margin-top:5px;color:#999;font-size:12px}

  .presale-indicator {
  display: inline-block;
  *display: inline;
  *zoom: 1;
  width: 24px;
  height: 15px;
  line-height: 15px;
  background: url(https://img3.doubanio.com/f/book/1679c65572eac1371f9872807199dea6e55a7f06/pics/book/cart/presale_text.gif) no-repeat;
  text-indent: -9999px;
  vertical-align: middle;
  *vertical-align: 0px;
  _vertical-align: 2px;
  margin-left: 0.5em;
}

</style>

  <div class="gray_ad buyinfo" id="buyinfo">
      <div class="buyinfo-printed" id="buyinfo-printed">
        

          
  

  <h2>
    <span>当前版本有售</span>
      &nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·

  </h2>


          
  <ul class="bs current-version-list">
      
      <li>
        

            
              <div class="cell price-btn-wrapper">
                <div class="vendor-name">
                  <a target="_blank" href="https://book.douban.com/link2/?lowest=1700&amp;pre=0&amp;vendor=jingdong&amp;srcpage=subject&amp;price=2960&amp;pos=1&amp;url=https%3A%2F%2Funion-click.jd.com%2Fjdc%3Fe%3D%26p%3DAyIHZRtYFAcXBFIZWR0yEgdUHF8cAxQ3EUQDS10iXhBeGlcJDBkNXg9JHUlSSkkFSRwSB1QcXxwDFBgMXgdIMkAAIklTXHdvZDdtAkgHUmY8QTBwYkQLWStcEAUUAFwYXCUCFgZRH2slAhM3B3WDo7TKouQHj7%252BNx4%252FCK1olAhYAVhxcFgUQBFMcUiUFIt7VrI23kMuZ387TgjIiN2UraxUyETcKXwZIMhM%253D%26t%3DW1dCFFlQCxxKQgFHRE5XDVULR0UVAhMAURJaEx1LQglG&amp;cntvendor=2&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=jingdong&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                    <span>京东商城</span>
                  </a>
                </div>

                <div class="cell impression_track_mod_buyinfo" data-track="https://frodo.douban.com/rohirrim/tracking/impression?subject_id=10594787&amp;isbn=7544258971&amp;source=jingdong&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject">
                  <div class="cell price-wrapper">
                    <a target="_blank" href="https://book.douban.com/link2/?lowest=1700&amp;pre=0&amp;vendor=jingdong&amp;srcpage=subject&amp;price=2960&amp;pos=1&amp;url=https%3A%2F%2Funion-click.jd.com%2Fjdc%3Fe%3D%26p%3DAyIHZRtYFAcXBFIZWR0yEgdUHF8cAxQ3EUQDS10iXhBeGlcJDBkNXg9JHUlSSkkFSRwSB1QcXxwDFBgMXgdIMkAAIklTXHdvZDdtAkgHUmY8QTBwYkQLWStcEAUUAFwYXCUCFgZRH2slAhM3B3WDo7TKouQHj7%252BNx4%252FCK1olAhYAVhxcFgUQBFMcUiUFIt7VrI23kMuZ387TgjIiN2UraxUyETcKXwZIMhM%253D%26t%3DW1dCFFlQCxxKQgFHRE5XDVULR0UVAhMAURJaEx1LQglG&amp;cntvendor=2&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=jingdong&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                      
                      <span class="buylink-price ">
                        
                        29.60元
                      </span>
                    </a>

                  </div>

                  <div class="cell">
                    <a href="https://book.douban.com/link2/?lowest=1700&amp;pre=0&amp;vendor=jingdong&amp;srcpage=subject&amp;price=2960&amp;pos=1&amp;url=https%3A%2F%2Funion-click.jd.com%2Fjdc%3Fe%3D%26p%3DAyIHZRtYFAcXBFIZWR0yEgdUHF8cAxQ3EUQDS10iXhBeGlcJDBkNXg9JHUlSSkkFSRwSB1QcXxwDFBgMXgdIMkAAIklTXHdvZDdtAkgHUmY8QTBwYkQLWStcEAUUAFwYXCUCFgZRH2slAhM3B3WDo7TKouQHj7%252BNx4%252FCK1olAhYAVhxcFgUQBFMcUiUFIt7VrI23kMuZ387TgjIiN2UraxUyETcKXwZIMhM%253D%26t%3DW1dCFFlQCxxKQgFHRE5XDVULR0UVAhMAURJaEx1LQglG&amp;cntvendor=2&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" class="buy-book-btn paper-book-btn" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=jingdong&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                      
                      <span>购买纸质书</span>
                    </a>
                  </div>
                </div>
              </div>

                
  

      </li>
      
      <li>
        

            
              <div class="cell price-btn-wrapper">
                <div class="vendor-name">
                  <a target="_blank" href="https://book.douban.com/link2/?lowest=1700&amp;pre=0&amp;vendor=kongfz&amp;srcpage=subject&amp;price=1700&amp;pos=2&amp;url=http%3A%2F%2Fitem.kongfz.com%2Fbook%2F44327900.html%3Fpush_type%3D5%26min_price%3D17.00%26utm_source%3D101004009001&amp;cntvendor=2&amp;link2key=6ff09ab291&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                    <span>孔网</span>
                  </a>
                </div>

                <div class="cell impression_track_mod_buyinfo" data-track="https://frodo.douban.com/rohirrim/tracking/impression?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject">
                  <div class="cell price-wrapper">
                    <a target="_blank" href="https://book.douban.com/link2/?lowest=1700&amp;pre=0&amp;vendor=kongfz&amp;srcpage=subject&amp;price=1700&amp;pos=1&amp;url=http%3A%2F%2Fitem.kongfz.com%2Fbook%2F44327900.html%3Fpush_type%3D5%26min_price%3D17.00%26utm_source%3D101004009001&amp;cntvendor=2&amp;link2key=6ff09ab291&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                      
                      <span class="buylink-price ">
                        
                        17.00元起
                      </span>
                    </a>

                  </div>

                  <div class="cell">
                    <a href="https://book.douban.com/link2/?lowest=1700&amp;pre=0&amp;vendor=kongfz&amp;srcpage=subject&amp;price=1700&amp;pos=1&amp;url=http%3A%2F%2Fitem.kongfz.com%2Fbook%2F44327900.html%3Fpush_type%3D5%26min_price%3D17.00%26utm_source%3D101004009001&amp;cntvendor=2&amp;link2key=6ff09ab291&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" class="buy-book-btn paper-book-btn" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                      
                      <span>购买纸质书</span>
                    </a>
                  </div>
                </div>
              </div>

                
  

      </li>
  </ul>


      </div>
      
  <div class="add2cartContainer no-border">
    
  <span class="add2cartWidget ll">
      <a class="j  add2cart a_show_login" href="https://www.douban.com/accounts/passport/login?source=book" rel="nofollow">
        <span>+ 加入购书单</span></a>
  </span>
  

  </div>

  </div>

<script>
  $(document).ready(function() {
    $('.impression_track_mod_buyinfo').each(function(i, item) {
      if (item) {
        var itmbUrl = $(item)[0]['dataset']['track']
        reportTrack(itmbUrl)
      }
    })
  })

  function track(url) {
    reportTrack(url)
  }

  function reportTrack(url) {
    if (!url) { return false }
    $.ajax({ url: url, dataType: 'text/html' })
  }
</script>













  






  

<style type="text/css" media="screen">
  .version_works h2{margin-bottom:20px}.version_works li.mb8{display:flex;flex-direction:column;padding-bottom:10px}.version_works li.mb8:last-child{padding-bottom:0;margin-bottom:0}.version_works .meta-wrapper{justify-content:space-between;display:flex}.version_works .meta-wrapper .count span{padding-right:10px}.version_works .buyinfo{display:flex;flex-direction:column;margin-top:10px;background:rgba(255,255,255,0.7)}.version_works .buyinfo li{padding:10px}.version_works .current-version-list{margin-bottom:0}.version_works .secondhand-books-list{margin:0}

</style>

<div class="gray_ad version_works">
  
  

  <h2>
    <span>这本书的其他版本 </span>
      &nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·
      <span class="pl">&nbsp;(
          <a href="https://book.douban.com/works/1011252">全部31</a>
        ) </span>

  </h2>


  <ul>
      
      <li class="mb8 pl">
        <div class="meta-wrapper">
          <div class="meta">
            <a href="https://book.douban.com/subject/3466923/">Penguin Classics （2007）</a>
            <div class="count">
              <span>9.1分</span> 452人读过
            </div>
          </div>

            <div class="fold-btn">
               <a href="javascript:void(0);"><span>展开</span>有售 (2)</a>
            </div>
        </div>

        <ul class="buyinfo" style="display: none;">
          
  <ul class="bs current-version-list">
      
      <li>
        

            
              <div class="cell price-btn-wrapper">
                <div class="vendor-name">
                  <a target="_blank" href="https://book.douban.com/link2/?pre=0&amp;vendor=jingdong&amp;srcpage=subject&amp;price=7980&amp;pos=1&amp;url=http%3A%2F%2Funion-click.jd.com%2Fjdc%3Fe%3D%26p%3DAyIOZR5aEQISA1AYUyUCGgRWElMTAiJDCkMFSjJLQhBaUAscSkIBR0ROVw1VC0dFFQoRBFwTXRUdS0IJRmtxW0Z6N28fUWJkT1xwHhZ%252Fc1kvUztlDh43Uh5cEwUbBFIrWxEDFgNlK1sUMkBpja3tzaejG4Gx1MCKhTdUK1sQCxsDUBhYFQYVAVwrXCXbkrCDucnMnJjS3YxrJTIi%26t%3DW1dCFBBFC1pXUwkEAEAdQFkJBVsdAREOXR1bCltXWwg%253D&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=jingdong&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                    <span>京东商城</span>
                  </a>
                </div>

                <div class="cell impression_track_manually" data-track="https://frodo.douban.com/rohirrim/tracking/impression?subject_id=10594787&amp;isbn=7544258971&amp;source=jingdong&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject">
                  <div class="cell price-wrapper">
                    <a target="_blank" href="https://book.douban.com/link2/?pre=0&amp;vendor=jingdong&amp;srcpage=subject&amp;price=7980&amp;pos=1&amp;url=http%3A%2F%2Funion-click.jd.com%2Fjdc%3Fe%3D%26p%3DAyIOZR5aEQISA1AYUyUCGgRWElMTAiJDCkMFSjJLQhBaUAscSkIBR0ROVw1VC0dFFQoRBFwTXRUdS0IJRmtxW0Z6N28fUWJkT1xwHhZ%252Fc1kvUztlDh43Uh5cEwUbBFIrWxEDFgNlK1sUMkBpja3tzaejG4Gx1MCKhTdUK1sQCxsDUBhYFQYVAVwrXCXbkrCDucnMnJjS3YxrJTIi%26t%3DW1dCFBBFC1pXUwkEAEAdQFkJBVsdAREOXR1bCltXWwg%253D&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=jingdong&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                      
                      <span class="buylink-price ">
                        
                        79.80元
                      </span>
                    </a>

                  </div>

                  <div class="cell">
                    <a href="https://book.douban.com/link2/?pre=0&amp;vendor=jingdong&amp;srcpage=subject&amp;price=7980&amp;pos=1&amp;url=http%3A%2F%2Funion-click.jd.com%2Fjdc%3Fe%3D%26p%3DAyIOZR5aEQISA1AYUyUCGgRWElMTAiJDCkMFSjJLQhBaUAscSkIBR0ROVw1VC0dFFQoRBFwTXRUdS0IJRmtxW0Z6N28fUWJkT1xwHhZ%252Fc1kvUztlDh43Uh5cEwUbBFIrWxEDFgNlK1sUMkBpja3tzaejG4Gx1MCKhTdUK1sQCxsDUBhYFQYVAVwrXCXbkrCDucnMnJjS3YxrJTIi%26t%3DW1dCFBBFC1pXUwkEAEAdQFkJBVsdAREOXR1bCltXWwg%253D&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" class="buy-book-btn paper-book-btn" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=jingdong&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                      
                      <span>购买纸质书</span>
                    </a>
                  </div>
                </div>
              </div>

                
  

      </li>
      
      <li>
        

            
              <div class="cell price-btn-wrapper">
                <div class="vendor-name">
                  <a target="_blank" href="https://book.douban.com/link2/?pre=0&amp;vendor=kongfz&amp;srcpage=subject&amp;price=2980&amp;pos=2&amp;url=http%3A%2F%2Fitem.kongfz.com%2Fbook%2F65449955.html%3Fpush_type%3D5%26min_price%3D29.80%26utm_source%3D101004009001&amp;link2key=6ff09ab291&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                    <span>孔网</span>
                  </a>
                </div>

                <div class="cell impression_track_manually" data-track="https://frodo.douban.com/rohirrim/tracking/impression?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject">
                  <div class="cell price-wrapper">
                    <a target="_blank" href="https://book.douban.com/link2/?pre=0&amp;vendor=kongfz&amp;srcpage=subject&amp;price=2980&amp;pos=1&amp;url=http%3A%2F%2Fitem.kongfz.com%2Fbook%2F65449955.html%3Fpush_type%3D5%26min_price%3D29.80%26utm_source%3D101004009001&amp;link2key=6ff09ab291&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                      
                      <span class="buylink-price ">
                        
                        29.80元起
                      </span>
                    </a>

                  </div>

                  <div class="cell">
                    <a href="https://book.douban.com/link2/?pre=0&amp;vendor=kongfz&amp;srcpage=subject&amp;price=2980&amp;pos=1&amp;url=http%3A%2F%2Fitem.kongfz.com%2Fbook%2F65449955.html%3Fpush_type%3D5%26min_price%3D29.80%26utm_source%3D101004009001&amp;link2key=6ff09ab291&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" class="buy-book-btn paper-book-btn" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                      
                      <span>购买纸质书</span>
                    </a>
                  </div>
                </div>
              </div>

                
  

      </li>
  </ul>


          
  <ul class="secondhand-books-list bs">
</ul>

        </ul>
      </li>
      
      <li class="mb8 pl">
        <div class="meta-wrapper">
          <div class="meta">
            <a href="https://book.douban.com/subject/26374532/">南海出版公司 （2015）</a>
            <div class="count">
              <span>8.9分</span> 7465人读过
            </div>
          </div>

            <div class="fold-btn">
               <a href="javascript:void(0);"><span>展开</span>有售 (3)</a>
            </div>
        </div>

        <ul class="buyinfo" style="display: none;">
          
  <ul class="bs current-version-list">
      
      <li>
        

            
              <div class="cell price-btn-wrapper">
                <div class="vendor-name">
                  <a target="_blank" href="https://book.douban.com/link2/?pre=0&amp;vendor=jingdong&amp;srcpage=subject&amp;price=3710&amp;pos=1&amp;url=http%3A%2F%2Funion-click.jd.com%2Fjdc%3Fe%3D%26p%3DAyIOZR5aEQISA1AYUyUCEgFUHVwWAiJDCkMFSjJLQhBaUAscSkIBR0ROVw1VC0dFFQIUBlMcWBUdS0IJRmtGW3lhVlI8T2AVTwB6O3dBbHwJRltDDh43Uh5cEwUbBFIrWxEDFgNlK1sUMkBpja3tzaejG4Gx1MCKhTdUK1sQCxsDUBteHQoSAVYrXCXbkrCDucnMnJjS3YxrJTIi%26t%3DW1dCFBBFC1pXUwkEAEAdQFkJBVsVBBMBUhhbCltXWwg%253D&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=jingdong&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                    <span>京东商城</span>
                  </a>
                </div>

                <div class="cell impression_track_manually" data-track="https://frodo.douban.com/rohirrim/tracking/impression?subject_id=10594787&amp;isbn=7544258971&amp;source=jingdong&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject">
                  <div class="cell price-wrapper">
                    <a target="_blank" href="https://book.douban.com/link2/?pre=0&amp;vendor=jingdong&amp;srcpage=subject&amp;price=3710&amp;pos=1&amp;url=http%3A%2F%2Funion-click.jd.com%2Fjdc%3Fe%3D%26p%3DAyIOZR5aEQISA1AYUyUCEgFUHVwWAiJDCkMFSjJLQhBaUAscSkIBR0ROVw1VC0dFFQIUBlMcWBUdS0IJRmtGW3lhVlI8T2AVTwB6O3dBbHwJRltDDh43Uh5cEwUbBFIrWxEDFgNlK1sUMkBpja3tzaejG4Gx1MCKhTdUK1sQCxsDUBteHQoSAVYrXCXbkrCDucnMnJjS3YxrJTIi%26t%3DW1dCFBBFC1pXUwkEAEAdQFkJBVsVBBMBUhhbCltXWwg%253D&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=jingdong&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                      
                      <span class="buylink-price ">
                        
                        37.10元
                      </span>
                    </a>

                  </div>

                  <div class="cell">
                    <a href="https://book.douban.com/link2/?pre=0&amp;vendor=jingdong&amp;srcpage=subject&amp;price=3710&amp;pos=1&amp;url=http%3A%2F%2Funion-click.jd.com%2Fjdc%3Fe%3D%26p%3DAyIOZR5aEQISA1AYUyUCEgFUHVwWAiJDCkMFSjJLQhBaUAscSkIBR0ROVw1VC0dFFQIUBlMcWBUdS0IJRmtGW3lhVlI8T2AVTwB6O3dBbHwJRltDDh43Uh5cEwUbBFIrWxEDFgNlK1sUMkBpja3tzaejG4Gx1MCKhTdUK1sQCxsDUBteHQoSAVYrXCXbkrCDucnMnJjS3YxrJTIi%26t%3DW1dCFBBFC1pXUwkEAEAdQFkJBVsVBBMBUhhbCltXWwg%253D&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" class="buy-book-btn paper-book-btn" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=jingdong&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                      
                      <span>购买纸质书</span>
                    </a>
                  </div>
                </div>
              </div>

                
  

      </li>
      
      <li>
        

            
              <div class="cell price-btn-wrapper">
                <div class="vendor-name">
                  <a target="_blank" href="https://book.douban.com/link2/?pre=0&amp;vendor=bookschina&amp;srcpage=subject&amp;price=3470&amp;pos=2&amp;url=http%3A%2F%2Fwww.bookschina.com%2Funion%2Fubook.asp%3Fadservice%3D354872%26tourl%3Dhttp%3A%2F%2Fwww.bookschina.com%2F6907929.htm&amp;link2key=6ff09ab291&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=bookschina&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                    <span>中国图书网</span>
                  </a>
                </div>

                <div class="cell impression_track_manually" data-track="https://frodo.douban.com/rohirrim/tracking/impression?subject_id=10594787&amp;isbn=7544258971&amp;source=bookschina&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject">
                  <div class="cell price-wrapper">
                    <a target="_blank" href="https://book.douban.com/link2/?pre=0&amp;vendor=bookschina&amp;srcpage=subject&amp;price=3470&amp;pos=1&amp;url=http%3A%2F%2Fwww.bookschina.com%2Funion%2Fubook.asp%3Fadservice%3D354872%26tourl%3Dhttp%3A%2F%2Fwww.bookschina.com%2F6907929.htm&amp;link2key=6ff09ab291&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=bookschina&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                      
                      <span class="buylink-price ">
                        
                        34.70元
                      </span>
                    </a>

                  </div>

                  <div class="cell">
                    <a href="https://book.douban.com/link2/?pre=0&amp;vendor=bookschina&amp;srcpage=subject&amp;price=3470&amp;pos=1&amp;url=http%3A%2F%2Fwww.bookschina.com%2Funion%2Fubook.asp%3Fadservice%3D354872%26tourl%3Dhttp%3A%2F%2Fwww.bookschina.com%2F6907929.htm&amp;link2key=6ff09ab291&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" class="buy-book-btn paper-book-btn" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=bookschina&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                      
                      <span>购买纸质书</span>
                    </a>
                  </div>
                </div>
              </div>

                
  

      </li>
      
      <li>
        

            
              <div class="cell price-btn-wrapper">
                <div class="vendor-name">
                  <a target="_blank" href="https://book.douban.com/link2/?pre=0&amp;vendor=kongfz&amp;srcpage=subject&amp;price=1000&amp;pos=3&amp;url=http%3A%2F%2Fitem.kongfz.com%2Fbook%2F35458201.html%3Fpush_type%3D5%26min_price%3D10.00%26utm_source%3D101004009001&amp;link2key=6ff09ab291&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                    <span>孔网</span>
                  </a>
                </div>

                <div class="cell impression_track_manually" data-track="https://frodo.douban.com/rohirrim/tracking/impression?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject">
                  <div class="cell price-wrapper">
                    <a target="_blank" href="https://book.douban.com/link2/?pre=0&amp;vendor=kongfz&amp;srcpage=subject&amp;price=1000&amp;pos=1&amp;url=http%3A%2F%2Fitem.kongfz.com%2Fbook%2F35458201.html%3Fpush_type%3D5%26min_price%3D10.00%26utm_source%3D101004009001&amp;link2key=6ff09ab291&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                      
                      <span class="buylink-price ">
                        
                        10.00元起
                      </span>
                    </a>

                  </div>

                  <div class="cell">
                    <a href="https://book.douban.com/link2/?pre=0&amp;vendor=kongfz&amp;srcpage=subject&amp;price=1000&amp;pos=1&amp;url=http%3A%2F%2Fitem.kongfz.com%2Fbook%2F35458201.html%3Fpush_type%3D5%26min_price%3D10.00%26utm_source%3D101004009001&amp;link2key=6ff09ab291&amp;srcsubj=10594787&amp;type=bkbuy&amp;subject=10594787" class="buy-book-btn paper-book-btn" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
                      
                      <span>购买纸质书</span>
                    </a>
                  </div>
                </div>
              </div>

                
  

      </li>
  </ul>


          
  <ul class="secondhand-books-list bs">
</ul>

        </ul>
      </li>
      
      <li class="mb8 pl">
        <div class="meta-wrapper">
          <div class="meta">
            <a href="https://book.douban.com/subject/1292227/">敦煌文艺出版社 （2000）</a>
            <div class="count">
              <span>8.8分</span> 9531人读过
            </div>
          </div>

            <div class="fold-btn">
               <a href="javascript:void(0);"><span>展开</span>有售 (1)</a>
            </div>
        </div>

        <ul class="buyinfo" style="display: none;">
          
  <ul class="bs current-version-list">
  </ul>


          
  <ul class="secondhand-books-list bs">
    
    <li class="impression_track_manually" data-track="https://frodo.douban.com/rohirrim/tracking/impression?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject">

        <div class="flex-wrap">
          <div class="cell vendor-name">
            <a target="_blank" href="http://item.kongfz.com/book/24891077.html?push_type=3&amp;min_price=3.00&amp;utm_source=101004009001" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
              <span>孔网</span>
            </a>
          </div>

          <div class="cell">
            <a target="_blank" href="http://item.kongfz.com/book/24891077.html?push_type=3&amp;min_price=3.00&amp;utm_source=101004009001" class="buylink-price" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
              <span>
                3.00元起
              </span>
            </a>
          </div>
          <div class="cell">
            <a target="_blank" href="http://item.kongfz.com/book/24891077.html?push_type=3&amp;min_price=3.00&amp;utm_source=101004009001" class="buy-book-btn secondhand-book-btn" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
              <span>购买二手书</span>
            </a>
          </div>
        </div>

        
  



    </li>
</ul>

        </ul>
      </li>
      
      <li class="mb8 pl">
        <div class="meta-wrapper">
          <div class="meta">
            <a href="https://book.douban.com/subject/1467368/">黑龙江人民出版社 （1987）</a>
            <div class="count">
              <span>9.0分</span> 2001人读过
            </div>
          </div>

            <div class="fold-btn">
               <a href="javascript:void(0);"><span>展开</span>有售 (1)</a>
            </div>
        </div>

        <ul class="buyinfo" style="display: none;">
          
  <ul class="bs current-version-list">
  </ul>


          
  <ul class="secondhand-books-list bs">
    
    <li class="impression_track_manually" data-track="https://frodo.douban.com/rohirrim/tracking/impression?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject">

        <div class="flex-wrap">
          <div class="cell vendor-name">
            <a target="_blank" href="http://item.kongfz.com/book/26054662.html?push_type=2&amp;min_price=19.80&amp;utm_source=101004009001" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
              <span>孔网</span>
            </a>
          </div>

          <div class="cell">
            <a target="_blank" href="http://item.kongfz.com/book/26054662.html?push_type=2&amp;min_price=19.80&amp;utm_source=101004009001" class="buylink-price" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
              <span>
                19.80元起
              </span>
            </a>
          </div>
          <div class="cell">
            <a target="_blank" href="http://item.kongfz.com/book/26054662.html?push_type=2&amp;min_price=19.80&amp;utm_source=101004009001" class="buy-book-btn secondhand-book-btn" onclick="track('https://frodo.douban.com/rohirrim/tracking/click?subject_id=10594787&amp;isbn=7544258971&amp;source=kongfz&amp;user_id=&amp;bid=SJ02sjYB7c8&amp;platform=pc&amp;location=vendor_subject')">
              <span>购买二手书</span>
            </a>
          </div>
        </div>

        
  



    </li>
</ul>

        </ul>
      </li>
  </ul>
</div>

<script>
  $(document).ready(function() {
    $('.fold-btn a').click(function() {
      var $btn = $(this).find('span');
      var $target = $(this).parents('.meta-wrapper').eq(0).next('.buyinfo');
      if ($target.is(':visible')) {
        $target.css('display', 'none');
        $btn.text('展开');
      } else {
        $target.css('display', 'flex');
        $btn.text('收起');

        // track
        if (!($target.attr('data-exposed'))) {
          $target.find('.impression_track_manually').each(function(i, item) {
            if (item) {
              var itmbUrl = $(item)[0]['dataset']['track']
              reportTrack(itmbUrl)
            }
          })
        }

        $target.attr('data-exposed', true);
      }
    })
  })
</script>


  <div id="dale_book_subject_top_middle" ad-status="loaded"></div>
  





<div class="gray_ad" id="borrowinfo">
  
  

  <h2>
    <span>在哪儿借这本书</span>
      &nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·

  </h2>


  <ul class="bs more-after">
      
      <li style="border: none">
        <a href="https://www.douban.com/link2/?url=http%3A%2F%2Fwww.bplisn.net.cn%2FBookSearch.aspx%3FISBN%3D978-7-5442-5897-5&amp;subject=7544258971&amp;type=borrow&amp;library=10007&amp;link2key=6ff09ab291" target="_blank">北京市公共图书馆(3)</a>
      </li>
      
      <li style="border: none">
        <a href="https://www.douban.com/link2/?url=http%3A%2F%2Fipac.library.sh.cn%2Fipac20%2Fipac.jsp%3Faspect%3Dbasic_search%26profile%3Dsl%26index%3DISBN%26term%3D9787544258975&amp;subject=7544258971&amp;type=borrow&amp;library=10012&amp;link2key=6ff09ab291" target="_blank">上海市中心图书馆(64)</a>
      </li>
      
      <li style="border: none">
        <a href="https://www.douban.com/link2/?url=http%3A%2F%2F58.154.49.3%3A8080%2Fopac%2Fopenlink.php%3FhistoryCount%3D1%26strText%3D978-7-5442-5897-5%26doctype%3DALL%26strSearchType%3Disbn%26match_flag%3Dforward%26displaypg%3D20%26sort%3DCATA_DATE%26orderby%3Ddesc%26showmode%3Dlist%26location%3DALL&amp;subject=7544258971&amp;type=borrow&amp;library=10010&amp;link2key=6ff09ab291" target="_blank">沈阳师范大学图书馆</a>
      </li>
      
      <li style="border: none">
        <a href="https://www.douban.com/link2/?url=http%3A%2F%2Fopac3.wzlib.cn%2Fopac%2Fsearch%3FsearchWay%3Disbn%26q%3D978-7-5442-5897-5%26booktype%3D%26marcformat%3D%26sortWay%3Dscore%26sortOrder%3Ddesc%26startPubdate%3D%26endPubdate%3D%26rows%3D10&amp;subject=7544258971&amp;type=borrow&amp;library=10020&amp;link2key=6ff09ab291" target="_blank">温州市图书馆</a>
      </li>
      
      <li style="border: none">
        <a href="https://www.douban.com/link2/?url=http%3A%2F%2Flaohu.sx.zj.cn%2Fweb%2Fwebsearch.asp%3FWCI%3DBookListH%26taab%3D7544258971&amp;subject=7544258971&amp;type=borrow&amp;library=10025&amp;link2key=6ff09ab291" target="_blank">绍兴图书馆</a>
      </li>
  </ul>
  <div class="clearfix"></div>
  <!--<div class="ft pl">-->
    <!--<a class="rr"  href="https://book.douban.com/library_invitation">&gt; 图书馆合作</a>-->
    <!--找不到你需要的图书馆？-->
  <!--</div>-->
</div>

  



    
  

  <h2>
    <span>以下书单推荐</span>
      &nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·&nbsp;·
      <span class="pl">&nbsp;(
          <a href="https://book.douban.com/subject/10594787/doulists">全部</a>
        ) </span>

  </h2>


    <div id="db-doulist-section" class="indent">
      <ul class="bs">
            
              <li class=""><a class="" href="https://www.douban.com/doulist/1757387/" target="_blank">评分9.0~9.7(1000+人参与评价)</a>
              </li>
            
              <li class=""><a class="" href="https://www.douban.com/doulist/13525/" target="_blank">电子书</a>
              </li>
            
              <li class=""><a class="" href="https://www.douban.com/doulist/436036/" target="_blank">你好好睡，我好好读书。等你醒了我们聊柏拉图。</a>
              </li>
            
              <li class=""><a class="" href="https://www.douban.com/doulist/2455132/" target="_blank">也许能改变你生活态度的书</a>
              </li>
            
              <li class=""><a class="" href="https://www.douban.com/doulist/43430373/" target="_blank">豆瓣高分书2700本：千人打分不低于8分</a>
              </li>
      </ul>
    </div>

  <div id="dale_book_subject_middle_mini" ad-status="appended" data-sell-type="COMPLEMENT" data-type="DoubanRender" data-version="3.2.28"><div style="position: relative; margin: 0px 0px 20px; display: block; width: 300px; height: 100px; overflow: hidden;"><iframe src="https://ad.doubanio.com" sandbox="allow-forms allow-scripts allow-same-origin allow-popups allow-top-navigation-by-user-activation" safe-mode="custom" width="300" height="100" frameborder="0" scrolling="no" name="dale_book_subject_middle_mini_frame" id="dale_book_subject_middle_mini_frame" style="overflow: hidden; display: block;"></iframe><div style="line-height: 1; text-align: center; background-color: rgba(0, 0, 0, 0.3); font-size: 12px; position: absolute; right: 0px; bottom: 0px; padding: 4px; color: rgb(255, 255, 255);">广告</div></div></div>
  






  <h2>谁读这本书?</h2>
  <div class="indent" id="collector">

    

<div class="">
    
    <div class="ll"><a href="https://www.douban.com/people/mushan12135/"><img src="https://img3.doubanio.com/icon/u152835149-11.jpg" class="pil" alt="是木杉啊！"></a></div>
    <div style="padding-left:60px"><a class="" href="https://www.douban.com/people/mushan12135/">是木杉啊！</a><br>
      <div class="pl ll">          4分钟前          想读      </div>

      <br>

    </div>
    <div class="clear"></div><br>
    <div class="ul" style="margin-bottom:12px;"></div>
</div>
<div class="">
    
    <div class="ll"><a href="https://www.douban.com/people/70599137/"><img src="https://img9.doubanio.com/icon/u70599137-5.jpg" class="pil" alt="晨曦之露"></a></div>
    <div style="padding-left:60px"><a class="" href="https://www.douban.com/people/70599137/">晨曦之露</a><br>
      <div class="pl ll">          6分钟前          想读      </div>

      <br>

    </div>
    <div class="clear"></div><br>
    <div class="ul" style="margin-bottom:12px;"></div>
</div>
<div class="">
    
    <div class="ll"><a href="https://www.douban.com/people/204436602/"><img src="https://img9.doubanio.com/icon/u204436602-5.jpg" class="pil" alt="和泉一织"></a></div>
    <div style="padding-left:60px"><a class="" href="https://www.douban.com/people/204436602/">和泉一织</a><br>
      <div class="pl ll">          6分钟前          读过      </div>

      <br>

    </div>
    <div class="clear"></div><br>
    <div class="ul" style="margin-bottom:12px;"></div>
</div>
<div class="">
    
    <div class="ll"><a href="https://www.douban.com/people/Seashellee/"><img src="https://img2.doubanio.com/icon/u216036191-2.jpg" class="pil" alt="竹子"></a></div>
    <div style="padding-left:60px"><a class="" href="https://www.douban.com/people/Seashellee/">竹子</a><br>
      <div class="pl ll">          7分钟前          读过      </div>

      <br>

    </div>
    <div class="clear"></div><br>
    <div class="ul" style="margin-bottom:12px;"></div>
</div>


        <p class="pl">&gt; <a href="https://book.douban.com/subject/10594787/comments?status=N">29456人在读</a></p>
        <p class="pl">&gt; <a href="https://book.douban.com/subject/10594787/comments?status=P">259079人读过</a></p>
        <p class="pl">&gt; <a href="https://book.douban.com/subject/10594787/comments?status=F">250205人想读</a></p>

  </div>





  
<!-- douban ad begin -->
<div id="dale_book_subject_middle_right" ad-status="loaded" data-sell-type="RTB" data-type="YoudaoRender" data-version="3.2.28"></div>
<script type="text/javascript">
    (function (global) {
        if(!document.getElementsByClassName) {
            document.getElementsByClassName = function(className) {
                return this.querySelectorAll("." + className);
            };
            Element.prototype.getElementsByClassName = document.getElementsByClassName;

        }
        var articles = global.document.getElementsByClassName('article'),
            asides = global.document.getElementsByClassName('aside');

        if (articles.length > 0 && asides.length > 0 && articles[0].offsetHeight >= asides[0].offsetHeight) {
            (global.DoubanAdSlots = global.DoubanAdSlots || []).push('dale_book_subject_middle_right');
        }
    })(this);
</script>
<!-- douban ad end -->

  





  

  <h2>二手市场</h2>
  <div class="indent">
    <ul class="bs">
    <li class="">
        <a class=" " href="https://book.douban.com/subject/10594787/offers">105本二手书欲转让</a>
          <span class="">
            (0.01
                至 30.00元)
          </span>
      </li>
      <li>
          <a class="rr j a_show_login" href="https://www.douban.com/register?reason=secondhand-offer&amp;cat=book"><span>在豆瓣转让</span></a>

        有250205人想读,手里有一本闲着?
      </li>
      <li>
        <div class="rr">
          <a href="https://shoushu.kongfz.com/?utm_source=douban&amp;utm_medium=new_offer">孔网上门收书</a>
        </div>
        转让给其他二手平台？
      </li>
    </ul>
  </div>

  
<p class="pl">订阅关于霍乱时期的爱情的评论: <br><span class="feed">
    <a href="https://book.douban.com/feed/subject/10594787/reviews"> feed: rss 2.0</a></span></p>


      </div>
      <div class="extra">
        
  
<!-- douban ad begin -->
<div id="dale_book_subject_bottom_super_banner" ad-status="loaded" data-sell-type="RTB" data-type="YoudaoRender" data-version="3.2.28"></div>
<script type="text/javascript">
    (function (global) {
        var body = global.document.body,
            html = global.document.documentElement;

        var height = Math.max(body.scrollHeight, body.offsetHeight, html.clientHeight, html.scrollHeight, html.offsetHeight);
        if (height >= 2000) {
            (global.DoubanAdSlots = global.DoubanAdSlots || []).push('dale_book_subject_bottom_super_banner');
        }
    })(this);
</script>
<!-- douban ad end -->


      </div>
    </div>
  </div>`

	result := introRe.FindString(str)

	fmt.Println(result)
}

