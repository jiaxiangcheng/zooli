<body>
    <link rel="stylesheet" href="/static/semantic-ui/dist/semantic.min.css"></link>
    <script src="/static/dist/semantic-ui/semantic.min.js"></script>
    <form action="javascript:void(0);" class="ui form">
        <h2 class="title"><i class="user icon"></i>User Information</h2>
        <div class="field">
            <h3 id="username">{{ .UserInfo.Username}}</h3>
            <label id="role">{{ .UserInfo.Role.Name}}</label>
        </div>

        <div class="field">
            <label>Email</label>
            <input id="email" type="email" placeholder="Email" value="{{ .UserInfo.Email}}" required/>
        </div>
        <div class="field">
            <label>Name</label>
            <input id="name" type="text" placeholder="Name" value="{{ .UserInfo.Name}}" required/>
        </div>

        <input class="ui primary button" id="save" type="submit" value="Save" onclick="SaveUser();">
        <input class="ui button" id="cancel" type="submit" value="Cancel" onclick="CancelUser();">
    </form>
</body>


<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>

<script type="text/javascript">

    function SaveUser() {
        var UserName = document.getElementById("username").value;
        var PassWord = document.getElementById("password").value;
        console.log("new password = " + PassWord);

        $.ajax({
            async: false,
            type: "post",
            dataType: "json",
            url: "/users/saveUser",
            data: { "username": UserName, "password": PassWord },
            success: function (data) {
                console.log("user saved data = " + data);
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
    .save-cancel input {
        display:inline-block;
    }
</style>
