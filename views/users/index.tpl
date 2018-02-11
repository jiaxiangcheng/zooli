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
        <link rel="stylesheet" href="/static/layout/widgets.css">
		<link rel="stylesheet" href="/static/semantic-ui/dist/semantic.min.css"></link>
		<script src="/static/dist/semantic-ui/semantic.min.js"></script>

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
        <script type="text/javascript" src="https://www.gstatic.com/charts/loader.js"></script>

        <script type="text/javascript">
            google.charts.load('current', {'packages':['corechart']});
            google.charts.setOnLoadCallback(drawChart);

            function drawChart() {
              var data = google.visualization.arrayToDataTable([
              ['Users', 'State'],
              ['Active', 50],
              ['Inactive', 60],
            ]);

              var options = {'title':'Analytics', 'width':550, 'height':400};

              // Display the chart inside the <div> element with id="piechart"
              var chart = new google.visualization.PieChart(document.getElementById('piechart'));
              chart.draw(data, options);
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
			    <button class="ui right floated basic button" id="sidebutton"
                onclick="showsidebutton()">
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
						<a href="#ncomments">
							<span><i class="home icon"></i></span>
							<span>Home</span>
						</a>
					</li>
					<li>
						<a href="#act-users">
							<span><i class="bar chart icon"></i></span>
							<span>Analytics</span>
						</a>
					</li>
					<li>
						<a href="https://www.google.es">
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
				{{ template "common/flash.tpl" . }}

				<table id="table_users" class="ui celled table">
					<thead>
						<tr>
							<th class="center aligned">ID</th>
							<th class="center aligned">Username</th>
							<th class="center aligned">Password</th>
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
						<i class="add user icon"></i>Create user</button>
				</div>
                </br>

                <div class="widget" id="act-users">
                    <div class="title">
                        Number of active Users
                    </div>
                    <div class="chart"></div>
                </div>
                <div class="widget" id="nvisits">
					<div class="title">Number of visits</div>
					<div class="chart"></div>
				</div>
				<div class="widget" id="ncomments">
					<div class="title">Number of comments</div>
					<div class="chart"></div>
				</div>


			</div>
		</div>
	</body>
</html>
