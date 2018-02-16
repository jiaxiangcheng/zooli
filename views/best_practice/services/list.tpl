<h1 class="ui header">Services</h1>
{{template "best_practice/common/flash.tpl" .}}
<table class="ui single line striped collapsing table">
    <thead>
    <tr>
        <th class="center aligned">Name</th>
        <th class="center aligned"></th>
        <th></th>
    </tr>
    </thead>
    <tbody>
    {{ range .services }}
    <tr>
        <td class="center aligned">{{ .Name}}</td>
        <td class="center aligned">
            <button type="button"
                    class="ui basic button"
                    onclick="editServices('{{ .ID}}');">
                View
            </button>
        </td>
        {{if ne $.user.ID .ID}}
        <td class="center aligned">
            <button type="button"
                    class="ui negative button"
                    data-toggle="modal"
                    data-target=".bs-example-modal-sm"
                    onclick="deleteServices('{{ .ID}}');">
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
        onclick="newServices();"
        style="margin: 10px 10px">
    <i class="add service icon"></i>
    Create service
</button>



<script type="text/javascript">
    function newServices() {
        $.ajax({
            async: false,
            type: "get",
            url: "/services/new",
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
    /*
    function editServices(user_id) {
        $.ajax({
            async: false,
            type: "get",
            url: "/users/" + user_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
    function deleteServices(user_id) {
        $.ajax({
            async: false,
            type: "delete",
            url: "/users/" + user_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }*/
</script>
