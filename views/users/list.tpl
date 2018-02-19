
{{template "common/modal.tpl" .}}
<h1 class="ui header">Users</h1>
{{template "common/flash.tpl" .}}
<table class="ui single line striped collapsing table">
    <thead>
    <tr>
        <th class="center aligned">Username</th>
        <th class="center aligned">Role</th>
        <th class="center aligned">Name</th>
        <th class="center aligned">Email</th>
        <th class="center aligned"></th>
        <th></th>
    </tr>
    </thead>
    <tbody>
    {{ range .users }}
    <tr>
        <td class="center aligned">{{ .Username}}</td>
        <td class="center aligned">{{ .Role.Name}}</td>
        <td class="center aligned">{{ .Name}}</td>
        <td class="center aligned">{{ .Email}}</td>
        <td class="center aligned">
            <button type="button"
                    class="ui basic button"
                    onclick="editUser('{{ .ID}}');">
                View
            </button>
        </td>
        {{if ne $.user.ID .ID}}
        <td class="center aligned">
            <button type="button"
                    class="ui negative button"
                    onclick="openDeleteModal('{{ .ID}}');">
                Delete
            </button>
        </td>
        {{end}}
    </tr>
    {{ end }}
    </tbody>
</table>
<button type="button"
        title="View user"
        class="ui basic big button"
        onclick="newUser();"
        style="margin: 10px 10px">
    <i class="add user icon"></i>
    Create user
</button>



<script type="text/javascript">

    var userId;

    function newUser() {
        $.ajax({
            async: false,
            type: "get",
            url: "/users/new",
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
    function editUser(user_id) {
        $.ajax({
            async: false,
            type: "get",
            url: "/users/" + user_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
    function deleteUser(user_id) {
        $.ajax({
            async: false,
            type: "delete",
            url: "/users/" + user_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }

    function openDeleteModal(user_id) {
        $('#mini_modal .header').html("Alert");
        $('#mini_modal .content').html("Are you sure to delete user?");
        $('#mini_modal')
                .modal({
                    onApprove : function() {
                        deleteUser(user_id)
                    }
                })
                .modal('show');
    }
</script>