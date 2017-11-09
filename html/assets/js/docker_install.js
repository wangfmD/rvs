$(function() {
	$("#docker_install").hide();
    $("#docker_install_new").show();
   dockerVersion();
})


var docker_version=0;

function dockerVersion(){
	$.ajax({
	  url: "getdockerversion.do",
	  async: false,
	  dataType: "json",
	  type: "get",
	  contentType: "application/x-www-form-urlencoded",
	  success: function(res) {
		if(res.Status=="fail"&&res.Info!="请登录"){
			 window.location.href='activation_auto.html';
		}
		if(res.Status=="ok"&&res.Info!=""){
		  docker_version=1;
		}
	  }
	})
}
function login_install(){
	var r=confirm("请确认是单台全量部署，如错误需重新安装！");
	if (r==false) {
		return;
	}
    docker_install("0", true);
}
function management(){
	var r=confirm("请确认是二台部署（一号机管理平台），如错误需重新安装！");
	if (r==false) {
		return;
	}
	docker_install("1", true);
}

function resources(){
	var r=confirm("请确认是多台部署（二号机资源平台），如错误需重新安装！");
	if (r==false) {
		return;
	}
	docker_install("2", true);
}

function new_login_install(){
        var r=confirm("请确认是单台全量部署，如错误需重新安装！");
        if (r==false) {
                return;
        }
    docker_install("0", false);
}
function new_management(){
        var r=confirm("请确认是二台部署（一号机管理平台），如错误需重新安装！");
        if (r==false) {
                return;
        }
        docker_install("1", false);
}

function new_resources(){
        var r=confirm("请确认是多台部署（二号机流媒体），如错误需重新安装！");
        if (r==false) {
                return;
        }
        docker_install("2", false);
}

function docker_install(val, old){
	layer.load();
	if(docker_version==0){
		$.ajax({
		  url: "installdocker.do",
		  dataType: "json",
		  type: "get",
		  contentType: "application/x-www-form-urlencoded",
		  success: function(res) {
			if(res.Status=="fail"&&res.Info=="请登录"){
				window.location.href='activation_auto.html';
			}
			if(res.Status=="ok"){
			  alert(res.Msg);
			  window.location.href=(old?'':'new_')+'docker_port.html?value='+val;
			}else{
			  alert(res.Msg+":"+res.Info);
			}
			layer.closeAll('loading');
		}
	  })
	}else{
		alert("Docker已经安装");
		window.location.href=(old?'':'new_')+'docker_port.html?value='+val;
	}
	
}
