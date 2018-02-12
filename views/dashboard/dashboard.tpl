<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>

<script type="text/javascript">
	function showsidebutton() {
		var x = document.getElementById("si");
		if (x.style.display === "block") {
			x.style.display = "none";
		} else {
			x.style.display = "block";
		}
	}
</script>

<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
	    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    	<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>ZOOLI Dashboard</title>

        <link rel="stylesheet" href="/static/layout/dashboard.css">
		<link rel="stylesheet" href="/static/layout/content.css">
		<link rel="stylesheet" href="/static/layout/header.css">
		<link rel="stylesheet" href="/static/layout/nav.css">
		<link rel="stylesheet" href="/static/layout/side-nav.css">

		<link rel="stylesheet" href="/static/semantic-ui/dist/semantic.min.css"></link>
		<script src="/static/dist/semantic-ui/semantic.min.js"></script>

	</head>
	<body>
		<div class="header">
			{{template "users/components/header.html"}}
		</div>
		<div class="side-nav" id="si">
			<div class="logo">
				<i class="large database icon"></i>
				<span>Zooli</span>
			</div>
            {{template "users/components/navigation.html"}}
		</div>
        <div class="main-content">
            <div class="ui statistic">
                <div class="value">
                    12323
                </div>
                <div class="label">
                    Users
                </div>
            </div>
        </div>

	</body>
</html>

<style>
    a:hover {
        background-color: #87CEEB;
    }
</style>