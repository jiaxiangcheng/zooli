<div class="ui raised segment">
    <div class="ui error message"></div>
    {{template "common/flash.tpl" .}}
    <div class="field">
        <div class="two fields">
            <div class="field">
                <label>Username</label>
                <input name="username" value="{{.userForm.Username}}" type="text" placeholder="Username"/>
            </div>
            <div class="field">
                <label>Password</label>
                <input name="password" type="password" placeholder="Password"/>
            </div>
        </div>
    </div>
    <div class="field">
        <div class="fields">
            <div class="ten wide field">
                <label>Email</label>
                <input name="email" value="{{.userForm.Email}}" type="email" placeholder="Email"/>
            </div>
            <div class="six wide field">
                <label>Name</label>
                <input name="name" value="{{.userForm.Name}}" type="text" placeholder="Name"/>
            </div>
        </div>
    </div>
    <div class="two fields">
        <div class="field">
            <label>Role</label>
            <div class="field">
                <select name="role" class="ui fluid dropdown">
                    <option value="">Role</option>
                {{ range .roles }}
                    {{ if $.userForm }}
                        <option value="{{.ID}}" {{ if eq .ID $.userForm.RoleID}} selected {{end}}>{{.Name}}</option>
                    {{else}}
                        <option value="{{.ID}}">{{.Name}}</option>
                    {{end}}
                {{end}}
                </select>
            </div>
        </div>

        {{ if ne .userForm.Role.Name .roleAdmin}}
            <div class="field">
                <label>Store</label>
                <div class="field">
                    <select name="store" class="ui fluid dropdown">
                        <option value="">Store</option>
                    {{ range .stores }}
                        {{ if $.userForm }}
                            <option value="{{.ID}}" {{ if eq .ID $.userForm.StoreID}} selected {{end}}>{{.Name}}</option>
                        {{else}}
                            <option value="{{.ID}}">{{.Name}}</option>
                        {{end}}
                    {{end}}
                    </select>
                </div>
            </div>
        {{end}}
    </div>
</div>


<script>
    $(document)
            .ready(function() {
                $('.dropdown').dropdown();
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
                                        },
                                        {
                                            type   : 'minLength[6]',
                                            prompt : 'Your password must be at least {ruleValue} characters'
                                        }
                                    ]
                                },
                                email: {
                                    identifier  : 'email',
                                    optional   : true,
                                    rules: [
                                        {
                                            type   : 'email',
                                            prompt : 'Incorrect email format'
                                        }
                                    ]
                                },
                                name: {
                                    identifier  : 'name',
                                    rules: [
                                        {
                                            type   : 'empty',
                                            prompt : 'Please enter your name'
                                        }
                                    ]
                                },
                                role: {
                                    identifier  : 'role',
                                    rules: [
                                        {
                                            type   : 'empty',
                                            prompt : 'Please select a role'
                                        }
                                    ]
                                }
                            }
                        });
            });
</script>