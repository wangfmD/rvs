 $(function() {
	 
	 console.log(getParam().value);
    //循环所有输入框
    inputValue();
	$("#configTwo").hide();
    if(getParam().value=="0"){
		//外网IP和域名联动
		textinputValue.init({inputYumin:'#outerport',inputNetword:'#outerports'}).render();
		//配置单台机器
		getSingleConfig();
		$("#configOne").show();
		$("#configTwo").hide();
		$("#configThree").hide();
    }else if(getParam().value=="1"){
		//外网IP和域名联动
		textinputValue.init({inputYumin:'#outerportsTwo',inputNetword:'#outerportTwo'}).render();
		//配置多单台机器 一号机
		getMultioneConfig();
		$("#configOne").hide();
		$("#configTwo").show();
		$("#configThree").hide();
	}
	else if(getParam().value=="2"){
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
	  url: "getsingleconfig.do",
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
	if(ftpip==""){
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
	} else if (host==internalip || host==outerip || internalip==outerip) {
		layer.msg("内外网ip及域名不能重复");
		isWhether=false;
	}
	if(isWhether){
		$.ajax({
          url: "setsingleconfig.do",
          async: false,
          dataType: "json",
          type: "get",
		  data:{
			  ftpip:ftpip,
			  host:host,
			  internalip:internalip,
			  internalport:internalport,
			  outerip:outerip,
			  outerport:outerport
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
		url: "getmultioneconfig.do",
		async: false,
		dataType: "json",
		type: "get",
		contentType: "application/x-www-form-urlencoded",
		success: function(res) {
			if(res.Status=="fail"&&res.Info=="请登录"){
			 window.location.href='activation_auto.html';
		}
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
	  url: "getmultitwoconfig.do",
	  async: false,
	  dataType: "json",
	  type: "get",
	  contentType: "application/x-www-form-urlencoded",
	  success: function(res) {
		  if(res.Status=="fail"&&res.Info=="请登录"){
			 window.location.href='activation_auto.html';
		}
			var data=res.Info;
			$("#databaseip").val(data.databaseip);
			$("#databaseport").val(data.databaseport);
			$("#ftpipThree").val(data.ftpip);
			$("#password").val(data.password);
			$("#username").val(data.username);
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
	var resourceipTwo = $("#resourceipTwo").val();
	
	
	if (hostTwo==internalipTwo || hostTwo==outeripTwo || internalipTwo==outeripTwo) {
		layer.msg("内外网ip及域名不能重复");
		isWhether=false;
	} else if(hostTwo==''){
		layer.msg("请输入域名!");
		isWhether=false;
	}else if(internalipTwo==''){
		layer.msg("请输入内网IP!");
		isWhether=false;
	}else if(internalportTwo==''){
		layer.msg("请输入内网端口!");
		isWhether=false;
	}else if(outeripTwo==''){
		layer.msg("");
		isWhether=false;
	}else if(outerportTwo==''){
		layer.msg("请输入外网IP!");
		isWhether=false;
	}else if(outerportsTwo==''){
		layer.msg("请输入外网端口!");
		isWhether=false;
	}else if(resourceipTwo==''){
		layer.msg("请输入二号机器IP!");
		isWhether=false;
	}
	
	if(isWhether){
		$.ajax({
		  url: "setmultioneconfig.do",
		  async: false,
		  dataType: "json",
		  type: "get",
		  data:{
			  host:hostTwo,
			  internalip:internalipTwo,
			  internalport:internalportTwo,
			  outerip:outeripTwo,
			  outerport:outerportTwo,
			  resourceip:resourceipTwo
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
	var databaseip = $("#databaseip").val();
	var databaseport = $("#databaseport").val();
	var ftpipThree = $("#ftpipThree").val();
	var password = $("#password").val();
	var username = $("#username").val();
	
	if(databaseip==''){
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
	}
	if(isWhether){
			 $.ajax({
			  url: "setmultitwoconfig.do",
			  async: false,
			  dataType: "json",
			  type: "get",
			  data:{
				  databaseip:databaseip,
				  databaseport:databaseport,
				  ftpip:ftpipThree,
				  password:password,
				  username:username
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