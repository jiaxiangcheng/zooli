<body>
    <link rel="stylesheet" href="/static/semantic-ui/dist/semantic.min.css"></link>
    <script src="/static/dist/semantic-ui/semantic.min.js"></script>
    <form action="" class="ui form">
            <h2 class="title">New user</h2></br>
            <div class="field">
                <input id="username" type="text" placeholder="Username"/>
            </div>
            <div class="field">
                <input id="password" type="password" placeholder="Password"/ >
            </div>
            <button class="ui fluid blue button" onclick="RegisterUser()">
                <span style="width:20px;"></span>注册</p>
            </button>
    </form>
</body>

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
