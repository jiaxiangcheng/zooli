<h1 class="ui header">Services</h1>
{{template "common/flash.tpl" .}}
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
        <td class="center aligned">
            <button type="button"
                    class="ui negative button"
                    data-toggle="modal"
                    data-target=".bs-example-modal-sm"
                    onclick="deleteServices('{{ .ID}}');">
                Delete
            </button>
        </td>
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

    function editServices(service_id) {
        $.ajax({
            async: false,
            type: "get",
            url: "/services/" + service_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
    function deleteServices(service_id) {
        $.ajax({
            async: false,
            type: "delete",
            url: "/services/" + service_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
</script>
