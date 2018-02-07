<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>ZOOLI Dashboard</title>
    <link href='https://fonts.googleapis.com/css?family=Source+Sans+Pro:700, 600,500,400,300' rel='stylesheet' type='text/css'></link>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.5.0/css/font-awesome.min.css"></link>
    <link rel="stylesheet" href="http://code.ionicframework.com/ionicons/2.0.1/css/ionicons.min.css"></link>

    <link rel="stylesheet" href="/static/css/dashboard.css"></link>
    <script src="/static/js/dashboard.js"></script>
    <!--SEMANTIC UI-->
    <link rel="stylesheet" href="/static/semantic-ui/dist/semantic.min.css"></link>
    <script src="/static/dist/semantic-ui/semantic.min.js"></script>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>

    <script src="https://code.jquery.com/jquery-2.2.0.min.js"></script>
    <script src="https://code.highcharts.com/highcharts.js"></script>
    <script src="https://code.highcharts.com/modules/data.js"></script>


</head>
<body>
    <div class="header">
        <div class="logo">
            <i class="dashboard icon"></i>
            <span class="name">ZOOLI</span>
        </div>
        <a href="#" class="nav-trigger"><span></span></a>
    </div>
    <div class="side-nav">
        <div class="logo">
            <i class="large dashboard icon"></i>
            <span class="name">ZOOLI</span>
        </div>
        <nav>
            <ul>
                <li class="active">
                    <a href="/users">
                        <i class="user icon"></i>
                        <span class="name">Users</span>
                    </a>
                </li>
                <li>
                    <a href="/">
                        <i class="home icon"></i>
                        <span class="name">Home</span>
                    </a>
                </li>
                <li>
                    <a href="#">
                        <i class="bar chart icon"></i>
                        <span class="name">Analytics</span>
                    </a>
                </li>
                <li>
                    <a href="/login">
                        <i class="sign out icon"></i>
                        <span class="name">Sig out</span>
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
            <table class="ui celled table">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Username</th>
                        <th>Password</th>
                        <th></th>
                        <th></th>
                        <th></th>
                    </tr>
                </thead>
                <tbody>
                    {{ range .users }}
                    <tr>
                        <td>{{ .ID}}</td>
                        <td class="positive">{{ .Username}}</td>
                        <td>{{ .Password}}</td>
                        <td>
                            <button type="button"
                            title="View user"
                            class="ui tiny button"
                            onclick="getUser('{{ .ID}}');">
                            View
                            </button>
                        </td>
                        <td>
                            <button type="button"
                            title="Delete user"
                            class="ui tiny negative button"
                            data-toggle="modal"
                            data-target=".bs-example-modal-sm"
                            onclick="deleteUser('{{ .ID}}');">
                            Delete
                            </button>
                        </td>
                        <td>
                            <button
                            title="Update user"
                            class="ui tiny button"
                            onclick="updateUser('{{ .ID}}');">
                            Update
                            </button>
                        </td>
                    </tr>
                    {{ end }}
                    </tbody>
            </table>

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
    </script>

</div>
</div>
</body>
</html>
