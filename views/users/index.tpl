<table id="table_users" class="table table-striped">
    <thead>
        <tr>
            <th>username</th>
            <th>password</th>
            <th></th>
            <th></th>
            <th></th>
        </tr>
    </thead>
    <tbody>
        {{ range .users }}
        <tr>
            <td>{{ .Username}}</td>
            <td>{{ .Password}}</td>
            <td>
                <button type="button" 
                        title="View user" 
                        class="btn btn-default" 
                        onclick="LoadUser({{ .Username}})">
                    View
                </button>
            </td>
            <td>
                <button type="button" 
                        title="Delete user" 
                        class="btn btn-default" 
                        data-toggle="modal" 
                        data-target=".bs-example-modal-sm" 
                        onclick="DeleteUser({{ .Username}})">
                    Delete
                </button>
            </td>
            <td>
                <button type="button" 
                        title="Update user" 
                        class="btn btn-default btn-sm" 
                        onclick="UpdateUser({{ .Username}})">
                    Update
                </button>
            </td>
        </tr>
        {{ end }}
        </tbody>
    </table>
