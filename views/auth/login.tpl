<body>
    <div class="ui grid">
        <link rel="stylesheet" href="/static/semantic-ui/dist/semantic.min.css"></link>
        <script src="/static/dist/semantic-ui/semantic.min.js"></script>
        <div class="row">
            <div class="panel panel-default">
                <div class="title">
                    <div class="ui icon header">
                        <i class ="settings icon"></i>
                        <div class="content">
                            <label>LOGIN</label>
                        </div>
                    </div>
                </div>
                <div class="panel-body">
                    <form accept-charset="utf-8" role="form" class="ui form" method="POST" action='{{urlfor "LoginController.Login"}}'>
                        <div class="ui form">
                            <br/>
                            <div class="ui left icon input">
                                <input class="form-control" placeholder="User name" name="Username" required id="inputUsername" />
                                <i class="users icon"></i>
                            </div>
                        </div>
                        <div class="ui form">
                            <br/>
                            <div class="ui left icon input">
                                <input class="form-control" placeholder="Password" name="Password" type="password" value="" required pattern=".{4,}" title="Password title" id="inputPassword"  />
                                <i class="privacy icon"></i>
                            </div>
                        </div>
                        <div class="form-group">
                            <div class="col-sm-12">
                            <br/>
                                <input class="ui fluid large blue button" type="submit" value="Login">
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</body>

<style>
    body {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        min-height: 100vh;
        background-color: #DADADA;
    }
    .title {
        text-align: center;
    }
</style>
