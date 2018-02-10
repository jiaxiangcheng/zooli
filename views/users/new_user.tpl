<form action="">
        <div style="margin:5px 0px;">
            <input id="username" type="text" />
         </div>
        <div style="margin:5px 0px;">
            <input id="password" type="password"/ >
        </div>
        <button type="button"  onclick="RegisterUser()">
            <span style="width:20px;"></span>注册</p>
        </button>
</form>

<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>

<script type="text/javascript">

    function RegisterUser()
    {
        var UserName = document.getElementById("username").value;
        var PassWord = document.getElementById("password").value;
        $.ajax({
            async: false,
            type: "POST",
            url: "/users/register",
            data: { "username":UserName, "password":PassWord },
            success: function (data) {
                $('body').html(data); 
            }
        });
    }
    

</script>