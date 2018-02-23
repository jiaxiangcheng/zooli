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

        <div class="field">
            <label>Store (Optional)</label>
            <div class="ui fluid search selection dropdown" id="stores">
                <input name="stores" type="hidden">
                <i class="dropdown icon"></i>
                <span class="default text">Stores</span>
                <div class="menu">
                    <div class="ui icon search input">
                        <i class="search icon"></i>
                        <input type="text" placeholder="Search a store...">
                    </div>
                    <div class="scrolling menu">
                        {{ range $.stores }}
                            <div class="item" data-value="{{.ID}}">
                                <span class="text">{{.Name}}</span>
                            </div>
                        {{end}}
                    </div>
                </div>
            </div>
            
        </div>
    </div>
</div>


<script>
    $(document)
            .ready(function() {
                $('.dropdown').dropdown();
                {{if .userForm}}
                    $('#stores').dropdown('set selected', ["{{$.userForm.StoreID}}"]);
                {{end}}

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
