<div class="container">
    <div class="row">
        <div class="panel panel-default">
            <div class="panel-heading text-center">
                <h3 class="panel-title"><strong>Login</strong></h3>
            </div>
            <div class="panel-body">
                <form accept-charset="utf-8" role="form" class="form-horizontal" method="POST" action='{{urlfor "LoginController.Login"}}'>
                    <div class="form-group">
                        <label for="inputUsername" class="col-sm-3 control-label">User name</label>
                        <div class="col-sm-9">
                          <input class="form-control" placeholder="User name" name="Username" required id="inputUsername" />
                        </div>
                      </div>
                      <div class="form-group">
                        <label for="inputPassword" class="col-sm-3 control-label">Password</label>
                        <div class="col-sm-9">
			    		  <input class="form-control" placeholder="Password" name="Password" type="password" value="" required pattern=".{6,}" title="Password title" id="inputPassword"  />
                        </div>
                      </div>
                      <div class="form-group">
                        <div class="col-sm-12">
			    		    <input class="btn btn-lg btn-success btn-block" type="submit" value="Login">
                        </div>
                      </div>
                </form>
            </div>
        </div>
    </div>
</div>