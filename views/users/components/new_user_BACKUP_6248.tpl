
<body>
    <link rel="stylesheet" href="/static/semantic-ui/dist/semantic.min.css"></link>
    <script src="/static/dist/semantic-ui/semantic.min.js"></script>
    <form action="javascript:void(0);" class="ui form" onsubmit="InsertUser();">
        <h2 class="title">
            <i class="user icon"></i>
            New user
        </h2></br>
        <div class="field">
            <label>Username</label>
            <input id="username" type="text" placeholder="Username" required/>
        </div>
        <div class="field">
            <label>Password</label>
            <input id="password" type="password" placeholder="Password" required/>
        </div>
        <div class="field">
            <label>Email</label>
            <input id="email" name="email" type="email" placeholder="Email" required/>
        </div>
        <div class="field">
            <label>Name</label>
            <input id="name" type="text" placeholder="Name" required/>
        </div>
        <input class="ui fluid blue button" id="register" type="submit" value="注册">
    </form>
</body>

<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>

<script type="text/javascript">
    $('.ui.form')
        .form({
            fields: {
            email: {
                identifier: 'email',
                rules: [{
                    type   : 'email',
                    prompt : 'Please enter a valid email'
                }]
            }
        }
    });

    function InsertUser() {
        var userName = document.getElementById("username").value;
        var passWord = document.getElementById("password").value;
        var email = document.getElementById("email").value;
        var name = document.getElementById("name").value;
        $.ajax({
            async: false,
            type: "post",
            dataType: "json",
            url: "/users/existUserIf",
            data: { "username": userName },
            success: function (data) {
                if (data.existed) {
                    window.alert("user existed");
                }
                else if (!data.existed) {
                    $.ajax({
                        async: false,
                        type: "post",
                        url: "/users/insert",
                        data: { 
                            "username": userName, 
                            "password": passWord,
                            "email": email,
                            "name": name
                        },
                        success: function (data) {
                            $('body').html(data);
                        }
                    });
                }
            }
        });
    }
</script>

<<<<<<< HEAD:views/users/new_user.tpl
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
=======
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
>>>>>>> master:views/users/components/new_user.tpl
