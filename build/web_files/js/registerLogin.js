document.addEventListener('DOMContentLoaded',function(){


var username,email,password,confirmpass;

  $("#registerBtn").click(function(){
    username = $("#username").val();
    email = $("#mail").val();
    password = $("#pass").val();
    confirmpass = $("#confirmpassword").val();
    console.log(username+" "+email+" "+password+" "+confirmpass);

if(username == "" || email == "" || password == "" || confirmpass == "")
{

  alert("You need fill each field");
}
else {



    $.ajax({

        traditional: true,
        type:"POST",
        url: 'http://localhost:8080/auth/register',
        contentType: 'application/json',
        data: JSON.stringify({"username": username,"email": email,"password": password}),
        dataType: 'json',
        statusCode: {
    200: function (response) {
        alert("200 code")

    },
    201: function (response) {
      $("#signUpH1").text("Registration was successfull. Please check your email for activation.");

      $("#registerForm").css("display", "none");
      setTimeout(function() {
    location.reload();
}, 5000);

    },
    400: function (response) {
       alert('1');
       bootbox.alert('<span style="color:Red;">Error While Saving Outage Entry Please Check</span>', function () { });
    },
    404: function (response) {
       alert('1');
       bootbox.alert('<span style="color:Red;">Error While Saving Outage Entry Please Check</span>', function () { });
    }
 },


        success: function()
        {




        }

} );
}

  });






















}(document, window, 0));