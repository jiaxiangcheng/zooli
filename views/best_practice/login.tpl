{{template "best_practice/common/form/header.tpl"}}

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
                        <input type="text" name="username" placeholder="username">
                    </div>
                </div>
                <div class="field">
                    <div class="ui left icon input">
                        <i class="lock icon"></i>
                        <input type="password" name="password" placeholder="password">
                    </div>
                </div>
                <div class="ui fluid large blue submit button">Login</div>
            </div>

            <div class="ui error message"></div>

        </form>
    {{template "best_practice/common/flash.tpl" .}}
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