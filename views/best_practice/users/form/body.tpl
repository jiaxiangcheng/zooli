<div class="ui raised segment">
    <h2 class="title1" id="title">
        <i class="user icon"></i>
        New user
    </h2></br>
    <div class="field">
        <div class="two fields">
            <div class="field">
                <label>Username</label>
                <input name="username" value="{{.userForm.Username}}" type="text" placeholder="Username" required/>
            </div>
            <div class="field">
                <label>Password</label>
                <input type="password" placeholder="Password" required/>
            </div>
        </div>
    </div>
    <div class="field">
        <div class="fields">
            <div class="ten wide field">
                <label>Email</label>
                <input name="email" type="email" placeholder="Email" required/>
            </div>
            <div class="six wide field">
                <label>Name</label>
                <input name="name" type="text" placeholder="Name" required/>
            </div>
        </div>
    </div>
    <div class="six wide field">
        <label>Role</label>
        <select name="role" class="ui dropdown">
            {{ range .roles }}
            <option value="{{.ID}}">{{.Name}}</option>
            {{end}}
        </select>
    </div>
</div>