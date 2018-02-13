<link rel="stylesheet" href="/static/semantic-ui/dist/semantic.min.css"></link>
<script src="/static/dist/semantic-ui/semantic.min.js"></script>

<body>
	<form action="javascript:void(0);" class="ui form" onsubmit="InsertUser();" id="createform">
		<div class="ui raised segment">
			<h2 class="title1" id="title">
				<i class="user icon"></i>
				New user
			</h2></br>
			<div class="field">
				<div class="two fields">
					<div class="field">
						<label>Username</label>
						<input id="username" type="text" placeholder="Username" required/>
					</div>
					<div class="field">
						<label>Password</label>
						<input id="password" type="password" placeholder="Password" required/>
					</div>
				</div>
			</div>
			<div class="field">
				<div class="fields">
					<div class="ten wide field">
						<label>Email</label>
						<input id="email" name="email" type="email" placeholder="Email" required/>
					</div>
					<div class="six wide field">
						<label>Name</label>
						<input id="name" type="text" placeholder="Name" required/>
					</div>
				</div>
			</div>
			<div class="six wide field">
				<label>Role</label>
				<select class="ui dropdown">
				  <option value="1">Manager</option>
				  <option value="0">Administrator</option>
				</select>
			</div>
			<input class="ui fluid blue button" id="register" type="submit" value="注册">
		</div>
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

<style>
    #createform {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        min-height: 100vh;
    }
	#title {
		text-align: center;
	}
    a:hover {
        background-color: #87CEEB;
    }
</style>
