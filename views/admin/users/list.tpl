{{template "common/modal.tpl" .}}

<div class="row">
    <div class="column">

        <div class="ui segments">
            <div class="ui segment">
                <h1 class="ui header center aligned">{{i18n .Lang "users_table.title"}}</h1>
            </div>
            <div class="ui segment">
                {{template "common/flash.tpl" .}}
                <table class="ui compact selectable striped celled table tablet stackable" id="data_table" cellspacing="0" width="100%">
                    <thead>
                        <tr>
                            <th>{{i18n .Lang "table_attribute_names.username"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.role"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.name"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.email"}}</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ range .users }}
                        <tr>
                            <td>{{ .Username}}</td>
                            <td>{{ .Role.Name}}</td>
                            <td>{{ .Name}}</td>
                            <td>{{ .Email}}</td>
                            <td class="center aligned">
                                <i class="blue link pencil alternate icon" onclick="editUser('{{ .ID}}');"></i>
                                {{if ne $.user.ID .ID}}
                                <i class="red link trash alternate icon" onclick="openDeleteModal('{{ .ID}}');"></i>
                                {{end}}
                            </td>
                        </tr>
                        {{ end }}
                    </tbody>
                    <tfoot class="full-width">
                        <tr>
                            <th colspan="5">
                                <div class="ui right floated small primary labeled icon button" onclick="newUser();">
                                    <i class="user icon"></i> {{i18n .Lang "users_table.add_user"}}
                                </div>
                            </th>
                        </tr>
                    </tfoot>
                </table>
            </div>
        </div>
    </div>
</div>


<script type="text/javascript">
$(document).ready(function() {
    $('#data_table').DataTable({
        language: {
            "search": {{i18n .Lang "search input"}}
        },
        lengthChange: false,
        info: false
    }
);

});

function newUser() {
    $.ajax({
        async: false,
        type: "get",
        url: "/admin/users/new",
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}
function editUser(user_id) {
    $.ajax({
        async: false,
        type: "get",
        url: "/admin/users/" + user_id,
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}
function deleteUser(user_id) {
    $.ajax({
        async: false,
        type: "delete",
        url: "/admin/users/" + user_id,
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