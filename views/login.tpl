<style type="text/css">
    body {
        background-color: #DADADA;
    }
    body > .grid {
        height: 100%;
    }
    .image {
        margin-top: -100px;
    }
    .column {
        max-width: 450px;
    }
</style>


<div class="ui middle aligned center aligned grid">
    <div class="column">
        <h2 class="ui blue image header">
            <i class ="settings icon"></i>
            <div class="content">
                Zooli
            </div>
        </h2>

        <form class="ui large form" method="POST" action="/login">
            <div class="ui stacked segment">
                <div class="field">
                    <div class="ui left icon input">
                        <i class="user icon"></i>
                        <input type="text" name="username" value="{{ .userForm.Username }}" placeholder="{{i18n .Lang "forms.username"}}">
                    </div>
                </div>
                <div class="field">
                    <div class="ui left icon input">
                        <i class="lock icon"></i>
                        <input type="password" name="password" placeholder="{{i18n .Lang "forms.password"}}">
                    </div>
                </div>
                <div class="ui fluid large blue submit button">{{i18n .Lang "login.login"}}</div>
            </div>

            <div class="ui error message"></div>
            {{template "common/flash.tpl" .}}
        </form>

    </div>
</div>


<script>
    $(document)
            .ready(function() {
                $('.ui.form')
                        .form({
                            fields: {
                                username: {
                                    identifier  : 'username',
                                    rules: [
                                        {
                                            type   : 'empty',
                                            prompt : 'Please enter your username'
                                        }
                                    ]
                                },
                                password: {
                                    identifier  : 'password',
                                    rules: [
                                        {
                                            type   : 'empty',
                                            prompt : 'Please enter your password'
                                        }
                                    ]
                                }
                            }
                        });
            });
</script>
