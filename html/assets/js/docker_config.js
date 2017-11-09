 $(function() { 
    //文件列表数
    showConfigTree();
    $("#browser").treeview({
        toggle : function() {
            //console.log("%s was toggled.", $(this).find(">span").text());
        }
    });
    
})
//配置文件提交按钮
function login_config() {
    alert("docker配置成功");
}
function showConfigTree(){
    $.ajax({
          url: "getconfigtree.do",
          async: false,
          dataType: "json",
          type: "get",
          contentType: "application/x-www-form-urlencoded",
          success: function(res) {
            //console.log(recursive(res.Info))
			if(res.Status=="fail"&&res.Info=="请登录"){
				 window.location.href='activation_auto.html';
			}
            $("#browser").html(recursive(res.Info,""));

          }
      })
}

function showFileContent(src){
    var pathFile=src.getAttribute("path").substring(1);
     $("#pathFile").html(pathFile);
     $.ajax({
          url: "getfilecontent.do?filepath="+pathFile,
          async: false,
          dataType: "html",
          type: "get",
          contentType: "application/x-www-form-urlencoded",
          success: function(res) {
			if(res.Status=="fail"&&res.Info=="请登录"){
				 window.location.href='activation_auto.html';
			}
            $("#html-pre").html(res);
          }
      })
}

function recursive(obj,parentpath){
    var htmlConfig="";
    for(var i=0;i<obj.length;i++){
        var Info=obj[i];
        if(typeof(Info)=="string"){
            htmlConfig+= "<li><span class='file' path='"+parentpath+"/"+Info+"' onclick='showFileContent(this)'>"+Info+"</span></li>";
        }else{
            for(var info1 in Info){
                htmlConfig+= "<li class='closed'><span class='folder'>"+info1+"</span>";
                htmlConfig+="<ul>"
                htmlConfig+=recursive(Info[info1],parentpath+"/"+info1);
                htmlConfig+="</ul></li>"
            }
        }
    }
    return htmlConfig;
}

//下载配置文件
function downloadconfig(){
    window.location.href="downloadconfig.do";
          
}
//备份
function backupconfig(){
    $.ajax({
          url: "backupconfig.do",
          async: false,
          dataType: "json",
          type: "get",
		  data:{
			  safecode:$("#safecode").val()
		  },
          contentType: "application/x-www-form-urlencoded",
          success: function(res) {
            layer.msg(res.Msg);
          }
      })
}
//恢复
function recoveryconfig(){
    $.ajax({
          url: "recoveryconfig.do",
          async: false,
          dataType: "json",
          type: "get",
		  data:{
			  safecode:$("#safecode").val()
		  },
          contentType: "application/x-www-form-urlencoded",
          success: function(res) {
            layer.msg(res.Msg);
          }
      })
}
//提交
function replaceconfig(){
    layer.load();
    $.ajax({
          url: "replaceconfig.do",
          async: false,
          dataType: "json",
          type: "post",
		  data:{
			  safecode:$("#safecode").val()
		  },
          cache: false,
          processData: false,
          contentType: false,
          data: new FormData($('#myform')[0]),
          success: function(res) {
			console.log(res);
            layer.msg(res.Msg);
            layer.closeAll('loading');
          }
      })
}