<script type="text/javascript">
    function InsertUser()
    {
        var UserName = document.getElementById("username").value;
        var PassWord = document.getElementById("password").value;
        console.log("userName = " + UserName);
        console.log("password = " + PassWord);
        $.ajax({
            async: false,
            type: "post",
            dataType: "json",
            url: "/users/existUserIf",
            data: { "username": UserName },
            success: function (data) {
                if (data.existed) {
                    window.alert("user existed");
                }
                else if (!data.existed) {
                    $.ajax({
                        async: false,
                        type: "post",
                        url: "/users/insert",
                        data: { "username": UserName, "password": PassWord },
                        success: function (data) {
                            $('body').html(data);
                        }
                    });
                }
            }
        });
    }
</script>

<body>
    <form action="javascript:void(0);" class="ui form" onsubmit="InsertUser();" id="create-form">
        <h2 class="title1">
            <i class="user icon"></i>
            New user
        </h2></br>
        <div class="field">
            <input id="username" type="text" placeholder="Username" required/>
        </div>
        <div class="field">
            <input id="password" type="password" placeholder="Password" required/>
        </div>
        <input class="ui blue button" id="register" type="submit" value="注册">
    </form>
</body>

<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
