<script type="text/javascript">
    function RegisterUser()
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
                        url: "/users/register",
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
    <link rel="stylesheet" href="/static/semantic-ui/dist/semantic.min.css"></link>
    <script src="/static/dist/semantic-ui/semantic.min.js"></script>
    <form action="" class="ui form" onsubmit="RegisterUser();">
            <h2 class="title">
                <i class="user icon"></i>
                New user</h2></br>
            <div class="field">
                <input id="username" type="text" placeholder="Username" required/>
            </div>
            <div class="field">
                <input id="password" type="password" placeholder="Password" required/>
            </div>
            <input class="ui fluid blue button" id="register" type="submit" value="注册"></button>
    </form>
</body>

<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>

<style>
    body {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        min-height: 100vh;
    }
    .title {
        text-align: center;
    }
</style>
