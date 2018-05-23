$(document).ready(function() {


  $("#loginBtn").click(function(event) {
    var uname = $("#uname").val();
    if (uname == undefined || uname == "") {
      window.alter("请输入账号！");
      return false;
    }
    var pwd = $("#pwd").val();
    if (pwd == undefined || pwd == "") {
      alter("请输入密码！");
      return false;
    }
    return true;
  });



});
