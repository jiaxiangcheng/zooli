<h1 class="ui header" style="text-align:center;">Users</h1>
{{template "common/modal.tpl" .}}
{{template "common/flash.tpl" .}}
<table class="ui single line striped collapsing table" id="table_list" style="table-layout:fixed; width:100%;">
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
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Username}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Role.Name}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Name}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Email}}</td>
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

<div class="ui middle aligned center aligned grid">
    <button type="button"
            title="View user"
            id="create_btn"
            class="ui basic big button"
            onclick="newUser();"
            style="margin: 15px;">
        <i class="add user icon"></i>
        Create user
    </button>
</div>

<style>
    #table_list {
       margin-left:auto;
       margin-right:auto;
     }
</style>

<script type="text/javascript">
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
