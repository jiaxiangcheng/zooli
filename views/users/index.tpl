<table id="table_users" class="table table-striped">
    <thead>
        <tr>
            <th>ID</th>
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
                <td>{{ .ID}}</td>
            <td>{{ .Username}}</td>
            <td>{{ .Password}}</td>
            <td>
                <button type="button" 
                        title="View user" 
                        class="btn btn-default" 
                        onclick="getUser('{{ .ID}}');">
                    View
                </button>
            </td>
            <td>
                <button type="button" 
                        title="Delete user" 
                        class="btn btn-default" 
                        data-toggle="modal" 
                        data-target=".bs-example-modal-sm" 
                        onclick="deleteUser('{{ .ID}}');">
                    Delete
                </button>
            </td>
            <td>
                <button type="button" 
                        title="Update user" 
                        class="btn btn-default btn-sm" 
                        onclick="updateUser('{{ .ID}}');">
                    Update
                </button>
            </td>
        </tr>
        {{ end }}
        </tbody>
    </table>

    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>

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
