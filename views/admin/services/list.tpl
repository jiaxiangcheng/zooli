{{template "common/modal.tpl" .}}

<div class="ui divider"></div>

<div class="row">
    <div class="column">
        <div class="ui segments">
            <div class="ui segment">
                <h1 class="ui header" style="text-align:center;">{{i18n .Lang "services_table.title"}}</h1>
            </div>
            <div class="ui segment">
                {{template "common/flash.tpl" .}}
                <table class="ui compact selectable striped celled table tablet stackable" id="data_table" cellspacing="0" width="100%">
                    <thead>
                        <tr>
                            <th>{{i18n .Lang "table_attribute_names.name"}}</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ range .services }}
                        <tr>
                            <td style="overflow: hidden;text-overflow: ellipsis;">{{ .Name}}</td>
                            <td class="center aligned">
                                <i class="blue link pencil alternate icon" onclick="editService('{{ .ID}}');"></i>
                                <i class="red link trash alternate icon" onclick="openDeleteModal('{{ .ID}}');"></i>
                            </td>
                        </tr>
                        {{ end }}
                    </tbody>
                    <tfoot class="full-width">
                        <tr>
                            <th colspan="5">
                                <div class="ui right floated small primary labeled icon button" onclick="newService();">
                                    <i class="cubes icon"></i> {{i18n .Lang "services_table.add_service"}}
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
        //dom: 'Bfrtip',
        lengthChange: false,
        info: false
    }
);

});

function newService() {
    $.ajax({
        async: false,
        type: "get",
        url: "/admin/services/new",
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}

function editService(service_id) {
    $.ajax({
        async: false,
        type: "get",
        url: "/admin/services/" + service_id,
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}
function deleteService(service_id) {
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
