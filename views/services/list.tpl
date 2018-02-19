<h1 class="ui header" style="text-align:center;">Services</h1>
{{template "common/flash.tpl" .}}
<table class="ui single line striped collapsing table" id="table_list">
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
</br>

<div class="ui middle aligned center aligned grid">
    <button type="button"
            title="View user"
            class="ui basic big button"
            onclick="newServices();">
        <i class="add service icon"></i>
        Create service
    </button>
</div>

<style>
    #table_list {
       margin-left:auto;
       margin-right:auto;
     }
</style>

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
