<h1 class="ui header" style="text-align:center; margin-top: 20px;">Services</h1>
{{template "common/modal.tpl" .}}
{{template "common/flash.tpl" .}}
<div class="ui divider"></div>

<table class="ui single line striped collapsing table" id="services-table"
    style="margin-left:auto; margin-right:auto; table-layout:fixed; width:100%;">
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
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Name}}</td>
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
                    onclick="openDeleteModal('{{ .ID}}');">
                Delete
            </button>
        </td>
    </tr>
    {{ end }}
    </tbody>
</table>

<div class="ui middle aligned center aligned grid">
    <button type="button"
            title="View user"
            class="ui blue basic big button"
            onclick="newServices();"
            style="margin: 15px;">
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

    $(document).ready(function() {
        $('#services-table').DataTable();
    });

    function newServices() {
        $.ajax({
            async: false,
            type: "get",
            url: "/admin/services/new",
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }

    function editServices(service_id) {
        $.ajax({
            async: false,
            type: "get",
            url: "/admin/services/" + service_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
    function deleteServices(service_id) {
        $.ajax({
            async: false,
            type: "delete",
            url: "/admin/services/" + service_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }

    function openDeleteModal(service_id) {
        $('#mini_modal .header').html("Alert");
        $('#mini_modal .content').html("Are you sure to delete service?");
        $('#mini_modal')
                .modal({
                    onApprove : function() {
                        deleteServices(service_id)
                    }
                })
                .modal('show');
    }
</script>
