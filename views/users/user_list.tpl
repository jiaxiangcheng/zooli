<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>

<script type="text/javascript">
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
				$('body').html(data);
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

<link rel="stylesheet" href="/static/semantic-ui/dist/semantic.min.css"></link>
<script src="/static/dist/semantic-ui/semantic.min.js"></script>

<body>
	<div class="title">Users</div>
	<div class="main">
		<table id="table_users" class="ui celled table">
			<thead>
				<tr>
					<th class="center aligned">ID</th>
					<th class="center aligned">Username</th>
					<th class="center aligned">Password</th>
					<th class="center aligned">Email</th>
					<th class="center aligned">Name</th>
					<th class="center aligned"></th>
					<th></th>
				</tr>
			</thead>
			<tbody>
				{{ range .users }}
				<tr>
					<td class="center aligned">{{ .ID}}</td>
					<td class="center aligned">{{ .Username}}</td>
					<td class="center aligned">{{ .PasswordHash}}</td>
					<td class="center aligned">{{ .Email}}</td>
					<td class="center aligned">{{ .Name}}</td>
					<td class="center aligned">
						<button type="button"
								class="ui basic button"
								onclick="getUser('{{ .ID}}');">
							View
						</button>
					</td>
					<td class="center aligned">
						<button type="button"
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
				<i class="add user icon"></i>
			Create user
		</button>
		</div>
		</br>

	</div>
</body>

<style>
    a:hover {
        background-color: #87CEEB;
    }
</style>
