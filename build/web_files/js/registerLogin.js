document.addEventListener('DOMContentLoaded',function(){
var sessionID;
 
var username,email,password,confirmpass,firstname,lastname,selectedClass,openPort,oldPassword,confirmOldPassword,newPassword,confirmNewPassword;
var loginUsername, loginPassword;

var tokk,iduser,sessID;

    

  $("#registerBtn").click(function(){
    username = $("#username").val();
    firstname = $("#firstname").val();
    lastname = $("#lastname").val();
    selectedClass = $("#selectClass option:selected").val();  
    email = $("#mail").val();
    password = $("#pass").val();
    confirmpass = $("#confirmpassword").val();
    console.log(username+" "+email+" "+password+" "+confirmpass);

if(username == "" || email == "" || password == "" || confirmpass == ""|| firstname == ""|| lastname == "")
{

  alert("You need fill each field");
}
else {



    $.ajax({

        traditional: true,
        type:"POST",
        url: 'http://itsovy.sk:1122/auth/register',
        contentType: 'application/json',
        data: JSON.stringify({"username": username,"name":firstname,"surname":lastname,"email": email,"password": password,"class":selectedClass}),
        dataType: 'json',
        statusCode: {
    200: function (response) {
        

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


$("#loginBtn").click(function(){
  if(sessionID ==""||sessionID=="null"){
  localStorage.setItem("sessionID",generateSesssionId())
  console.log(sessionID)
      $("#h1Username").html(loginUsername);
}
else {
  sessionID = localStorage.getItem("sessionID");
  console.log(sessionID)
}
loginUsername = $("#loginUsername").val();
loginPassword = $("#loginPassword").val();
     
console.log(sessionID+" "+loginUsername+" "+loginPassword);

$.ajax({

    traditional: true,
    type:"POST",
    url: 'http://itsovy.sk:1122/auth/login',
    contentType: 'application/json',
    data: JSON.stringify({"sessionid":sessionID,"username": loginUsername,"password": loginPassword}),
    dataType: 'json',
    statusCode: {
200: function (response) {



    console.log(response.token)
      localStorage.setItem("usrname",loginUsername);
    localStorage.setItem("tokk",response.token);
    localStorage.setItem("iduser",response.iduser);
    localStorage.setItem("sessID",response.sessionid);
     sessID = sessionID = localStorage.getItem("sessID");
  tokk = localStorage.getItem("tokk");
  iduser = localStorage.getItem("iduser");
    alert(tokk+" "+sessID+" "+iduser);
    window.location.href = "loggedIn.html";
   
    

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






  });



$("#logoutBtn").click(function(){
  sessID = sessionID = localStorage.getItem("sessID");
  tokk = localStorage.getItem("tokk");
  iduser = localStorage.getItem("iduser");

  $.ajax({

      traditional: true,
      type:"POST",
      url: 'http://itsovy.sk:1122/auth/logout',
      contentType: 'application/json',
      data: JSON.stringify({"sessionid":sessID,"iduser":parseInt(iduser),"token":tokk}),
      dataType: 'json',
      statusCode: {
  200: function (response) {
window.location.href = "index.html";


  },
  201: function (response) {



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


});

    
    
    
$("#loggedIn").html('Logged as: '+localStorage.getItem("usrname"));

    $("#mysqlButton").click(function(){
        sessID = sessionID = localStorage.getItem("sessID");
        tokk = localStorage.getItem("tokk");
        iduser = localStorage.getItem("iduser");
      $("#h1Username").html(localStorage.getItem("usrname"));
        
          $.ajax({

      traditional: true,
      type:"POST",
      url: 'http://itsovy.sk:1122/linux/available',
      contentType: 'application/json',
      data: JSON.stringify({"auth":{"sessionid":sessID,"iduser":parseInt(iduser),"token":tokk}}),
      dataType: 'json',
      statusCode: {
  200: function (response) {
     console.log(response.ports);
        for(var i=0;i<response.ports.length;i++)
            {
      $("#portSelect").append("<option>"+response.ports[i].port+"</option>")

            }

  },
  201: function (response) {



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
        
        
        });
    
    $("#btnOpenPort").click(function(){
        openPort = $("#portSelect option:selected").val(); 
       

          $.ajax({

      traditional: true,
      type:"POST",
      url: 'http://itsovy.sk:1122/linux/openport',
      contentType: 'application/json',
      data: JSON.stringify({"auth":{"sessionid":sessID,"iduser":parseInt(iduser),"token":tokk},"port":parseInt(openPort)}),
      dataType: 'json',
      statusCode: {
  200: function (response) {
    
 myFunction();  
  },
  201: function (response) {



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
        
        
        
        
        
        
        
        
        
      });
    
    
     $("#btnChangePassword").click(function(){
        oldPassword =  $("#oldPassword").val();
        confirmOldPassword = $("#confirmOldPassword").val();
        newPassword = $("#newPassword").val();
        confirmNewPassword = $("#confirmNewPassword").val();
             
         console.log(oldPassword +" "+ newPassword);
       if(oldPassword != confirmOldPassword)
           {
               alert("Old passwords isn´t same");
               
               
           }
         
         else if(newPassword !=confirmNewPassword)
             {
                 alert("New passwords isn´t same");
             }
         
         else{
$.ajax({

      traditional: true,
      type:"POST",
      url: 'http://itsovy.sk:1122/auth/chpasswd',
      contentType: 'application/json',
      data: JSON.stringify({"auth":{"sessionid":sessID,"iduser":parseInt(iduser),"token":tokk},"oldpass":oldPassword,"newpass":newPassword}),
      dataType: 'json',
      statusCode: {
  200: function (response) {
    
 myFunction();  
  },
  201: function (response) {



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
    
    
      $("#btnClosePort").click(function(){
          
          
          
          $.ajax({

      traditional: true,
      type:"POST",
      url: 'http://itsovy.sk:1122/linux/myports',
      contentType: 'application/json',
      data: JSON.stringify({"auth":{"sessionid":sessID,"iduser":parseInt(iduser),"token":tokk}}),
      dataType: 'json',
      statusCode: {
  200: function (response) {
    for(var i=0;i<response.ports.length;i++){
  $("#openedPorts").append("<option>"+response.ports[i].port+"</option>")
    }
  },
  201: function (response) {



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
          
          
          
              });
      
        window.onload = function() {
    if (window.jQuery) {  
        $("#h1Username").html(localStorage.getItem("usrname"));
        
        
    } else {
        // jQuery is not loaded
        alert("Doesn't Work");
    }
}
    
     
    

        
}(document, window, 0));

    function myFunction() {
  // Get the snackbar DIV
  var x = document.getElementById("snackbar");

  // Add the "show" class to DIV
  x.className = "show";

  // After 3 seconds, remove the show class from DIV
  setTimeout(function(){ x.className = x.className.replace("show", ""); }, 3000);
}

  function generateSesssionId() {
    var text = "";
    var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";

    for (var i = 0; i < 16; i++)
      text += possible.charAt(Math.floor(Math.random() * possible.length));

    return text;
  }

  $(function(){

  $('li.dropdown > a').on('click',function(event){

    event.preventDefault()

    $(this).parent().find('ul').first().toggle(300);

    $(this).parent().siblings().find('ul').hide(200);

    //Hide menu when clicked outside
    $(this).parent().find('ul').mouseleave(function(){
      var thisUI = $(this);
      $('html').click(function(){
        thisUI.hide();
        $('html').unbind('click');
      });
    });


  });




});
