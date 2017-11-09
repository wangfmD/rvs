 $(function() {
	 
	 console.log(getParam().value);
    //循环所有输入框
    inputValue();
	$("#configTwo").hide();
    if(getParam().value=="0"){
		//外网IP和域名联动
		(Object.create(textinputValue)).init({inputYumin:'#outerport',inputNetword:'#outerports'}).render();
		(Object.create(textinputValue)).init({inputYumin:'#internalip',inputNetword:'#httpInterIp'}).render();
		(Object.create(textinputValue)).init({inputYumin:'#outerip',inputNetword:'#httpOuterIp'}).render();
		(Object.create(textinputValue)).init({inputYumin:'#host',inputNetword:'#httpHost'}).render();
		(Object.create(textinputValue)).init({inputYumin:'#httpOuterPort',inputNetword:'#httpHostPort'}).render();
		(Object.create(textinputValue)).init({inputYumin:'#internalip',inputNetword:'#streamInterIp'}).render();
		(Object.create(textinputValue)).init({inputYumin:'#outerip',inputNetword:'#streamOuterIp'}).render();
		(Object.create(textinputValue)).init({inputYumin:'#host',inputNetword:'#streamHost'}).render();
		(Object.create(textinputValue)).init({inputYumin:'#streamOuterPort',inputNetword:'#streamHostPort'}).render();

		//配置单台机器
		getSingleConfig();
		$("#configOne").show();
		$("#configTwo").hide();
		$("#configThree").hide();
    }else if(getParam().value=="1"){
		//外网IP和域名联动
		(Object.create(textinputValue)).init({inputYumin:'#outerportsTwo',inputNetword:'#outerportTwo'}).render();
		(Object.create(textinputValue)).init({inputYumin:'#streamip',inputNetword:'#resourceipTwo'}).render();

		//配置多单台机器 一号机
		getMultioneConfig();
		$("#configOne").hide();
		$("#configTwo").show();
		$("#configThree").hide();
	}
	else if(getParam().value=="2"){
		(Object.create(textinputValue)).init({inputYumin:'#streamOuterPort2',inputNetword:'#streamHostPort2'}).render();
		(Object.create(textinputValue)).init({inputYumin:'#httpOuterPort2',inputNetword:'#httpHostPort2'}).render();
		
		(Object.create(textinputValue)).init({inputYumin:'#httpInterIp2',inputNetword:'#streamInterIp2'}).render();
		(Object.create(textinputValue)).init({inputYumin:'#httpOuterIp2',inputNetword:'#streamOuterIp2'}).render();
		(Object.create(textinputValue)).init({inputYumin:'#httpHost2',inputNetword:'#streamHost2'}).render();
		//配置多单台机器 二号机
		getMultitWoconfig();
		$("#configOne").hide();
		$("#configTwo").hide();
		$("#configThree").show();
	}
    
  //隐藏全选
  $("#allInput").hide();
  //隐藏确定添加
  $("#determine").hide();
})

function showStreamInfo() {

	if ($('#checkInter').is(':checked')) {
		$("#interHttpDiv").show();
		$("#interStreamDiv").show();
	} else {
		$("#interHttpDiv").hide();
		$("#interStreamDiv").hide();
	}
	if ($('#checkOuter').is(':checked')) {
		$("#outerHttpDiv").show();
		$("#outerStreamDiv").show();
	} else {
		$("#outerHttpDiv").hide();
		$("#outerStreamDiv").hide();
	}
	if ($('#checkHost').is(':checked')) {
		$("#hostHttpDiv").show();
		$("#hostStreamDiv").show();
	} else {
		$("#hostHttpDiv").hide();
		$("#hostStreamDiv").hide();
	}
	
	if ($('#checkInter2').is(':checked')) {
		$("#interHttpDiv2").show();
		$("#interStreamDiv2").show();
	} else {
		$("#interHttpDiv2").hide();
		$("#interStreamDiv2").hide();
	}
	if ($('#checkOuter2').is(':checked')) {
		$("#outerHttpDiv2").show();
		$("#outerStreamDiv2").show();
	} else {
		$("#outerHttpDiv2").hide();
		$("#outerStreamDiv2").hide();
	}
	if ($('#checkHost2').is(':checked')) {
		$("#hostHttpDiv2").show();
		$("#hostStreamDiv2").show();
	} else {
		$("#hostHttpDiv2").hide();
		$("#hostStreamDiv2").hide();
	}
}

//http请求是否带参数
var getParam=function(){
	try{
		var url = window.location.href;
		var result = url.split("?")[1];
		var keyValue = result.split("&");
		var obj = {};
		for (var i = 0; i < keyValue.length; i++) {
			var item = keyValue[i].split("=");
			obj[item[0]] = item[1];
		}
		return obj;
	}catch(e){
		console.warn("There has no param value!");
	}
}
 
//配置单台机器
function getSingleConfig(){
	$.ajax({
	  url: "getsingleconfignew.do",
	  async: false,
	  dataType: "json",
	  type: "get",
	  contentType: "application/x-www-form-urlencoded",
	  success: function(res) {
			if(res.Status=="fail"&&res.Info=="请登录"){
				 window.location.href='activation_auto.html';
			}
			$("#ftpip").val(res.Info.ftpip);
			$("#host").val(res.Info.host);
			$("#internalip").val(res.Info.internalip);
			$("#internalport").val(res.Info.internalport);
			$("#outerip").val(res.Info.outerip);
			$("#outerport").val(res.Info.outerport);
			$("#outerports").val(res.Info.outerport);
			if (res.Info.http0 == "1")
				$("#checkInter").attr("checked",'true');
			if (res.Info.http1 == "1")
				$("#checkOuter").attr("checked",'true');
			if  (res.Info.http2 == "1")
				$("#checkHost").attr("checked",'true');
			$("#httpInterIp").val(res.Info.hhttp0);
			$("#httpInterPort").val(res.Info.phttp0);
			$("#httpOuterIp").val(res.Info.hhttp1);
			$("#httpOuterPort").val(res.Info.phttp1);
			$("#httpHost").val(res.Info.hhttp2);
			$("#httpHostPort").val(res.Info.phttp2);
			$("#streamInterIp").val(res.Info.hstreaming0);
			$("#streamInterPort").val(res.Info.pstreaming0);
			$("#streamOuterIp").val(res.Info.hstreaming1);
			$("#streamOuterPort").val(res.Info.pstreaming1);
			$("#streamHost").val(res.Info.hstreaming2);
			$("#streamHostPort").val(res.Info.pstreaming2);
			showStreamInfo();
		}
    })
}

function setSingleConfigSubbmit(){
	
	
	var isWhether=true;
	var ftpip = $("#ftpip").val();
	var host = $("#host").val();
	var internalip = $("#internalip").val();
	var internalport = $("#internalport").val();
	var outerip = $("#outerip").val();
	var outerport = $("#outerport").val();
	var outerports = $("#outerports").val();
	
	var http0 = $('#checkInter').is(':checked');
	var http1 = $('#checkOuter').is(':checked');
	var http2 = $('#checkHost').is(':checked');
	var hhttp0 = $("#httpInterIp").val();
	var phttp0= $("#httpInterPort").val();
	var hhttp1= $("#httpOuterIp").val();
	var phttp1= $("#httpOuterPort").val();
	var hhttp2= $("#httpHost").val();
	var phttp2= $("#httpHostPort").val();
	var hstreaming0= $("#streamInterIp").val();
	var pstreaming0= $("#streamInterPort").val();
	var hstreaming1= $("#streamOuterIp").val();
	var pstreaming1= $("#streamOuterPort").val();
	var hstreaming2= $("#streamHost").val();
	var pstreaming2= $("#streamHostPort").val();
	var shttp0 = http0?"1":"0";
	var shttp1 = http1?"1":"0";
	var shttp2 = http2?"1":"0";
	if ( !http0 && !http1 && !http2) {
		layer.msg("至少选择一个流媒体服务");
		return;
	} 
	if (http0 && hhttp0=="") {
		layer.msg("http服务内网ip不能为空");
		isWhether=false;
	} else if (http0 && phttp0=="") {
		layer.msg("http服务内网端口不能为空");
		isWhether=false;
	} else if (http0 && hstreaming0=="") {
		layer.msg("stream服务内网ip不能为空");
		isWhether=false;
	} else if (http0 && pstreaming0=="") {
		layer.msg("stream服务内网端口不能为空");
		isWhether=false;
	} else if (http1 && hhttp1=="") {
		layer.msg("http服务外网ip不能为空");
		isWhether=false;
	} else if (http1 && phttp1=="") {
		layer.msg("http服务外网端口不能为空");
		isWhether=false;
	} else if (http1 && hstreaming1=="") {
		layer.msg("stream服务外网ip不能为空");
		isWhether=false;
	} else if (http1 && pstreaming1=="") {
		layer.msg("stream服务外网端口不能为空");
		isWhether=false;
	} else if (http2 && hhttp2=="") {
		layer.msg("http服务域名不能为空");
		isWhether=false;
	} else if (http2 && phttp2=="") {
		layer.msg("http服务域名端口不能为空");
		isWhether=false;
	} else if (http2 && hstreaming2=="") {
		layer.msg("stream服务域名ip不能为空");
		isWhether=false;
	} else if (http2 && pstreaming2=="") {
		layer.msg("stream服务域名端口不能为空");
		isWhether=false;
	} else if (host==internalip || host==outerip || internalip==outerip) {
		layer.msg("内外网ip及域名不能重复");
		isWhether=false;
	} else if(ftpip==""){
		layer.msg("请输入Ftp地址!");
		isWhether=false;
	}else if(host==""){
		layer.msg("请输入域名!");
		isWhether=false;
	}else if(internalip==""){
		layer.msg("请输入内网IP!");
		isWhether=false;
	}else if(internalport==""){
		layer.msg("请输入内网IP端口!");
		isWhether=false;
	}else if(outerip==""){
		layer.msg("请输入外网IP!");
		isWhether=false;
	}else if(outerport==""){
		layer.msg("请输入外网IP端口!");
		isWhether=false;
	}
	if(isWhether){
		$.ajax({
          url: "setsingleconfignew.do",
          async: false,
          dataType: "json",
          type: "get",
		  data:{
			  ftpip:ftpip,
			  host:host,
			  internalip:internalip,
			  internalport:internalport,
			  outerip:outerip,
			  outerport:outerport,
			  hhttp0:hhttp0,
			  phttp0:phttp0,
			  hhttp1:hhttp1,
			  phttp1:phttp1,
			  hhttp2:hhttp2,
			  phttp2:phttp2,
			  hstreaming0:hstreaming0,
			  pstreaming0:pstreaming0,
			  hstreaming1:hstreaming1,
			  pstreaming1:pstreaming1,
			  hstreaming2:hstreaming2,
			  pstreaming2:pstreaming2,
			  http0:shttp0,
			  http1:shttp1,
			  http2:shttp2,
			  streaming0:shttp0,
			  streaming1:shttp1,
			  streaming2:shttp2
		  },
          contentType: "application/x-www-form-urlencoded",
          success: function(res) {
                layer.msg(res.Msg);
				//配置单台机器
				getSingleConfig();
          }
      })
	}
}

//配置多单台机器 一号机
function getMultioneConfig(){
	$.ajax({
		url: "getmultioneconfignew.do",
		async: false,
		dataType: "json",
		type: "get",
		contentType: "application/x-www-form-urlencoded",
		success: function(res) {
			if(res.Status=="fail"&&res.Info=="请登录"){
				 window.location.href='activation_auto.html';
			}
			$("#streamip").val(res.Info.streamip);
			$("#streamport").val(res.Info.streamport);
			$("#hostTwo").val(res.Info.host);
			$("#internalipTwo").val(res.Info.internalip);
			$("#internalportTwo").val(res.Info.internalport);
			$("#outeripTwo").val(res.Info.outerip);
			$("#outerportTwo").val(res.Info.outerport);
			$("#outerportsTwo").val(res.Info.outerport);
			$("#resourceipTwo").val(res.Info.resourceip);
		}
    })
}

//配置多单台机器 二号机
function getMultitWoconfig(){
	$.ajax({
	  url: "getmultitwoconfignew.do",
	  async: false,
	  dataType: "json",
	  type: "get",
	  contentType: "application/x-www-form-urlencoded",
	  success: function(res) {
			if(res.Status=="fail"&&res.Info=="请登录"){
				 window.location.href='activation_auto.html';
			}
			var data=res.Info;
			$("#authip").val(data.authip);
			$("#authport").val(data.authport);
			$("#databaseip").val(data.databaseip);
			$("#databaseport").val(data.databaseport);
			$("#ftpipThree").val(data.ftpip);
			$("#password").val(data.password);
			$("#username").val(data.username);
			if (res.Info.http0 == "1")
				$("#checkInter2").attr("checked",'true');
			if (res.Info.http1 == "1")
				$("#checkOuter2").attr("checked",'true');
			if  (res.Info.http2 == "1")
				$("#checkHost2").attr("checked",'true');
			$("#httpInterIp2").val(res.Info.hhttp0);
			$("#httpInterPort2").val(res.Info.phttp0);
			$("#httpOuterIp2").val(res.Info.hhttp1);
			$("#httpOuterPort2").val(res.Info.phttp1);
			$("#httpHost2").val(res.Info.hhttp2);
			$("#httpHostPort2").val(res.Info.phttp2);
			$("#streamInterIp2").val(res.Info.hstreaming0);
			$("#streamInterPort2").val(res.Info.pstreaming0);
			$("#streamOuterIp2").val(res.Info.hstreaming1);
			$("#streamOuterPort2").val(res.Info.pstreaming1);
			$("#streamHost2").val(res.Info.hstreaming2);
			$("#streamHostPort2").val(res.Info.pstreaming2);
			showStreamInfo();
	    }
    })
}

function setMultioneConfigSubbmit(){
	var isWhether=true;
	//---------一号机配置---------//
	var hostTwo = $("#hostTwo").val();
	var internalipTwo = $("#internalipTwo").val();
	var internalportTwo = $("#internalportTwo").val();
	var outeripTwo = $("#outeripTwo").val();
	var outerportTwo = $("#outerportTwo").val();
	var outerportsTwo = $("#outerportsTwo").val();
	var streamip = $("#streamip").val();
	var streamport = $("#streamport").val();
	var resourceip = $("#resourceipTwo").val();
	if (hostTwo==internalipTwo || hostTwo==outeripTwo || internalipTwo==outeripTwo) {
		layer.msg("内外网ip及域名不能重复");
		isWhether=false;
	} else if (resourceip=='') {
		layer.msg("请输入二号机地址!");
		isWhether=false;
	} else if(streamip==''){
		layer.msg("请输入流媒体服务端口!");
		isWhether=false;
	}else if(streamport==''){
		layer.msg("请输入流媒体服务ip!");
		isWhether=false;
	}else if(hostTwo==''){
		layer.msg("请输入域名!");
		isWhether=false;
	}else if(internalipTwo==''){
		layer.msg("请输入内网IP!");
		isWhether=false;
	}else if(internalportTwo==''){
		layer.msg("请输入内网端口!");
		isWhether=false;
	}else if(outeripTwo==''){
		layer.msg("请输入外网IP");
		isWhether=false;
	}else if(outerportTwo==''){
		layer.msg("请输入外网端口!");
		isWhether=false;
	}else if(outerportsTwo==''){
		layer.msg("请输入外网端口!");
		isWhether=false;
	}
	if(isWhether){
		$.ajax({
		  url: "setmultioneconfignew.do",
		  async: false,
		  dataType: "json",
		  type: "get",
		  data:{
			  resourceip:resourceip,
			  streamip:streamip,
			  streamport:streamport,
			  internalip:internalipTwo,
			  internalport:internalportTwo,
			  outerip:outeripTwo,
			  outerport:outerportTwo,
			  host:hostTwo
		  },
		  contentType: "application/x-www-form-urlencoded",
		  success: function(res) {
			  layer.msg(res.Msg);
			  //配置多单台机器 一号机
			  getMultioneConfig();
			  
		  }
	  })
		
	}
}

function setMultioneConfigSubbmits(){
	var isWhether=true;
	//--------二号机配置----------//
	var authip = $("#authip").val();
	var authport = $("#authport").val();
	var databaseip = $("#databaseip").val();
	var databaseport = $("#databaseport").val();
	var ftpipThree = $("#ftpipThree").val();
	var password = $("#password").val();
	var username = $("#username").val();
	
	var http0 = $('#checkInter2').is(':checked');
	var http1 = $('#checkOuter2').is(':checked');
	var http2 = $('#checkHost2').is(':checked');
	var hhttp0 = $("#httpInterIp2").val();
	var phttp0= $("#httpInterPort2").val();
	var hhttp1= $("#httpOuterIp2").val();
	var phttp1= $("#httpOuterPort2").val();
	var hhttp2= $("#httpHost2").val();
	var phttp2= $("#httpHostPort2").val();
	var hstreaming0= $("#streamInterIp2").val();
	var pstreaming0= $("#streamInterPort2").val();
	var hstreaming1= $("#streamOuterIp2").val();
	var pstreaming1= $("#streamOuterPort2").val();
	var hstreaming2= $("#streamHost2").val();
	var pstreaming2= $("#streamHostPort2").val();
	var shttp0 = http0?"1":"0";
	var shttp1 = http1?"1":"0";
	var shttp2 = http2?"1":"0";
	if ( !http0 && !http1 && !http2) {
		layer.msg("至少选择一个流媒体服务");
		return;
	} 
	if (http0 && hhttp0=="") {
		layer.msg("http服务内网ip不能为空");
		isWhether=false;
	} else if (http0 && phttp0=="") {
		layer.msg("http服务内网端口不能为空");
		isWhether=false;
	} else if (http0 && hstreaming0=="") {
		layer.msg("stream服务内网ip不能为空");
		isWhether=false;
	} else if (http0 && pstreaming0=="") {
		layer.msg("stream服务内网端口不能为空");
		isWhether=false;
	} else if (http1 && hhttp1=="") {
		layer.msg("http服务外网ip不能为空");
		isWhether=false;
	} else if (http1 && phttp1=="") {
		layer.msg("http服务外网端口不能为空");
		isWhether=false;
	} else if (http1 && hstreaming1=="") {
		layer.msg("stream服务外网ip不能为空");
		isWhether=false;
	} else if (http1 && pstreaming1=="") {
		layer.msg("stream服务外网端口不能为空");
		isWhether=false;
	} else if (http2 && hhttp2=="") {
		layer.msg("http服务域名不能为空");
		isWhether=false;
	} else if (http2 && phttp2=="") {
		layer.msg("http服务域名端口不能为空");
		isWhether=false;
	} else if (http2 && hstreaming2=="") {
		layer.msg("stream服务域名ip不能为空");
		isWhether=false;
	} else if (http2 && pstreaming2=="") {
		layer.msg("stream服务域名端口不能为空");
		isWhether=false;
	} else if(databaseip==''){
		layer.msg("请输入数据库Ip地址!");
		isWhether=false;
	}else if("#databaseport"==''){
		layer.msg("请输入数据库端口!");
		isWhether=false;
	}else if(ftpipThree==''){
		layer.msg("请输入FtpIpd地址!");
		isWhether=false;
	}else if(password==''){
		layer.msg("请输入MySql密码!");
		isWhether=false;
	}else if(username==''){
		layer.msg("请输入MySql用户名!");
		isWhether=false;
	}else if(authip==''){
		layer.msg("请输入认证服务ip!");
		isWhether=false;
	}else if(authport==''){
		layer.msg("请输入认证服务端口!");
		isWhether=false;
	}
	
	if(isWhether){
			 $.ajax({
			  url: "setmultitwoconfignew.do",
			  async: false,
			  dataType: "json",
			  type: "get",
			  data:{
				  authip:authip,
				  authport:authport,
				  databaseip:databaseip,
				  databaseport:databaseport,
				  ftpip:ftpipThree,
				  password:password,
				  username:username,
				  hhttp0:hhttp0,
			  phttp0:phttp0,
			  hhttp1:hhttp1,
			  phttp1:phttp1,
			  hhttp2:hhttp2,
			  phttp2:phttp2,
			  hstreaming0:hstreaming0,
			  pstreaming0:pstreaming0,
			  hstreaming1:hstreaming1,
			  pstreaming1:pstreaming1,
			  hstreaming2:hstreaming2,
			  pstreaming2:pstreaming2,
			  http0:shttp0,
			  http1:shttp1,
			  http2:shttp2,
			  streaming0:shttp0,
			  streaming1:shttp1,
			  streaming2:shttp2
			  },
			  contentType: "application/x-www-form-urlencoded",
			  success: function(res) {
				layer.msg(res.Msg);
				//配置多单台机器 二号机
				getMultitWoconfig();
			  }
		  })
	}
}

//垒加数字
var countInput=1;
//添加端口
function addPost(){
	//显示确定添加
	$("#determine").show();
	$("#addPostDocker").append('<div class="opera_item"> <div class="form-group row"> <label class="control-label col-sm-4">端口'+countInput+':</label> <div class="col-sm-6" style="padding-right: 0;"> <input id="pollcode" name="pollcode1" class="form-control" placeholder="请输入端口'+countInput+'"> </div> </div> </div>');
	countInput++;
}

function inputValue(){
	$.ajax({
	  url: "getallopenports.do",
	  async: false,
	  dataType: "json",
	  type: "get",
	  contentType: "application/x-www-form-urlencoded",
	  data: {
	  },
	  success: function(res) {
		var htmlInputButOther="";
		var htmlInputBut="";
		var lengthPort=res.Info.length/2;
		for(var i = 0 ; i<res.Info.length;i++){
			
			if(i<lengthPort){
				htmlInputBut+='<div class="opera_item"> <div class="form-group row"> <label class="control-label col-sm-3">端口'+(i+1)+':</label> <div class="col-sm-6" style="padding-right: 0;"> <input name="pollcode" class="form-control" placeholder="请输入端口'+(i+1)+'" value="'+res.Info[i]+'" disabled> </div> </div> </div>';
			}else if(i>=lengthPort){
				htmlInputButOther+='<div class="opera_item"> <div class="form-group row"> <label class="control-label col-sm-3">端口'+(i+1)+':</label> <div class="col-sm-6" style="padding-right: 0;"> <input name="pollcode" class="form-control" placeholder="请输入端口'+(i+1)+'" value="'+res.Info[i]+'" disabled>  </div> </div> </div>';
			}
			countInput++;
		}
		$("#myformInput").html(htmlInputBut);
		$("#myformInputOther").html(htmlInputButOther);
			
	  }
  })
}

//修改外面地址端口  域名也要一起边
var textinputValue={
    inputYumin:null,
    inputNetword:null,
    init:function(config){
        this.inputYumin=$(config.inputYumin);
        this.inputNetword=$(config.inputNetword);
        this.bind();
        return this;
    },
    bind:function(){
        var self=this;
        this.inputYumin.on('keyup',function(){
            self.render();
        });
    },
    getNum:function(){
        return this.inputYumin.val();
    },
    render:function(){
        var num=this.getNum();
        this.inputNetword.val(num);
    }
}


//下一步按钮
function login_init(){
	window.location.href="projectManager.html";
}

function openPort(){
	layer.load();
    var obj=document.getElementsByName('pollcode1'); //选择所有name="'test'"的对象，返回数组 
    //取到对象数组后，我们来循环检测它是不是被选中 
    var s='';
    for(var i=0; i<obj.length; i++){ 
       //如果选中，将value添加到变量s中  
		if(obj[i].value!=""){
			s+='&port='+obj[i].value
		}
    }
	if(s!=''){
		$.ajax({
          url: "openport.do",
          async: false,
          dataType: "json",
          type: "get",
          contentType: "application/x-www-form-urlencoded",
          data:s.substring(1),
          success: function(res) {
			if(res.Status=="ok"){
				layer.msg("添加成功！");
				inputValue();
			}
          }
		})
	}else{
		layer.msg("没有要添加的端口！");
	}
	layer.closeAll('loading');
}
