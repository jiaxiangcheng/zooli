<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>

<script type="text/javascript">

    function getUser(user_id) {
        $.ajax({
            async: false,
            type: "post",
            url: "/users/" + user_id,
            data: {
                id: user_id
            },
            success: function (data) {
                console.log("data = " + data);
                $('body').html(data);
            }
        });
    }

    function createUser() {
        $.ajax({
            async: false,
            type: "post",
            url: "/users/new",
            success: function (data) {
                $('body').html(data);
            }
        });
    }

</script>

<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
	    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    	<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>ZOOLI Dashboard</title>

		<link rel="stylesheet" href="/static/layout/main.css">
		<link rel="stylesheet" href="/static/layout/content.css">
		<link rel="stylesheet" href="/static/layout/header.css">
		<link rel="stylesheet" href="/static/layout/nav.css">
		<link rel="stylesheet" href="/static/layout/side-nav.css">
		<link rel="stylesheet" href="/static/semantic-ui/dist/semantic.min.css"></link>
		<script src="/static/dist/semantic-ui/semantic.min.js"></script>
		<script src="/static/css/dashboard.js"></script>

        <script type="text/javascript">
            function getUser(user_id) {
                console.log("user_id = " + user_id);
                $.ajax({
                    async: false,
                    type: "post",
                    url: "/users/" + user_id,
                    data: {
                        id: user_id
                    },
                    success: function (data) {
                        //$('body').html(data);
                    }
                });
            }
			function showsidebutton() {
				var x = document.getElementById("si");
			    if (x.style.display === "block") {
			        x.style.display = "none";
			    } else {
			        x.style.display = "block";
			    }
			}
        </script>

	</head>
	<body>
		<div class="header">
			<div class="logo">
				<i class="dashboard icon"></i>
				<span>Zooli</span>
			</div>
            <div class="sidebar">
			    <button class="ui right floated basic button" id="sidebutton" onclick="showsidebutton()">
					<i class="sidebar icon"></i></button>
            </div>
		</div>
		<div class="side-nav" id="si">
			<div class="logo">
				<i class="large dashboard icon"></i>
				<span>Zooli</span>
			</div>
			<nav>
				<ul>
					<li class="active">
						<a href="/users">
							<span><i class="user icon"></i></span>
							<span>Users</span>
						</a>
					</li>
					<li>
						<a href="/">
							<span><i class="home icon"></i></span>
							<span>Home</span>
						</a>
					</li>
					<li>
						<a href="#">
							<span><i class="bar chart icon"></i></span>
							<span>Analytics</span>
						</a>
					</li>
					<li>
						<a href="/login">
							<span><i class="sign out icon"></i></span>
							<span>Sig out</span>
						</a>
					</li>
				</ul>
			</nav>
		</div>
		<div class="main-content">
			<div class="title">
				Users
			</div>
			<div class="main">
				<table id="table_users" class="ui celled table">
					<thead>
						<tr>
							<th>ID</th>
							<th>username</th>
							<th>password</th>
							<th></th>
							<th></th>
						</tr>
					</thead>
					<tbody>
						{{ range .users }}
						<tr>
							<td>{{ .ID}}</td>
							<td>{{ .Username}}</td>
							<td>{{ .Password}}</td>
							<td>
								<button type="button"
										title="View user"
										class="ui primary button"
										onclick="getUser('{{ .ID}}');">
									View
								</button>
							</td>
							<td>
								<button type="button"
										title="Delete user"
										class="ui negative button"
										data-toggle="modal"
										data-target=".bs-example-modal-sm"
										onclick="deleteUser('{{ .ID}}');">
									Delete
								</button>
							</td>
						</tr>
						{{ end }}
					</tbody>
				</table>
				<button type="button"
						title="View user"
						class="ui basic big button"
						onclick="createUser();"
						style="margin: 10px 10px">
						<i class="add user icon"></i>Create user</button>
				</div>
			</div>
		</div>
	</body>
</html>
