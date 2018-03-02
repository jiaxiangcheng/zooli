{{template "common/modal.tpl" .}}

<div class="ui divider"></div>

<div class="row">
    <div class="column">

        <div class="ui segments">
            <div class="ui segment">
                <h1 class="ui header" style="text-align:center;">{{i18n .Lang "stores_table.title"}}</h1>
            </div>
            <div class="ui segment">
                {{template "common/flash.tpl" .}}
                <table class="ui compact selectable striped celled table tablet stackable" id="data_table" cellspacing="0" width="100%">
                    <thead>
                        <tr>
                            <th>{{i18n .Lang "table_attribute_names.name"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.address"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.phone"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.companies"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.managers"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.services"}}</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ range $i, $s := .stores }}
                        <tr>
                            <td style="overflow: hidden;text-overflow: ellipsis;">{{ .Name}}</td>
                            <td style="overflow: hidden;text-overflow: ellipsis;">{{ .Address}}</td>
                            <td style="overflow: hidden;text-overflow: ellipsis;">{{ .PhoneNumber}}</td>
                            <td style="overflow: hidden;text-overflow: ellipsis;">{{ .Company.Name}}</td>
                            <td style="overflow: hidden;text-overflow: ellipsis;">{{range .Managers}} <a class="ui olive label">{{.Name}}</a>{{end}}</td>
                            <td style="overflow: hidden;text-overflow: ellipsis;">{{range .Services}} <a class="ui blue label">{{.Name}}</a> {{end}}</td>
                            <td class="center aligned">
                                <i class="blue link pencil alternate icon" onclick="editStore('{{ .ID}}');"></i>
                                <i class="red link trash alternate icon" onclick="openDeleteModal('{{ .ID}}');"></i>
                            </td>
                        </tr>
                        {{ end }}
                    </tbody>
                    <tfoot class="full-width">
                        <tr>
                            <th colspan="5">
                                <div class="ui right floated small primary labeled icon button" onclick="newStore();">
                                    <i class="shopping bag icon"></i> {{i18n .Lang "stores_table.add_store"}}
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

$(document)
.ready(function() {
    /*$('.dropdown')
    .dropdown({
    action: function(text, value) {
    args = value.split(",");
    $.ajax({
    type: "post",
    url: "/admin/users/" + args[0] + "/assign",
    data: {"storeID": args[1]},
    success: function (data) {
    $('#main_content').html(data);
}
});
}
});*/
});


function newStore() {
    $.ajax({
        async: false,
        type: "get",
        url: "/admin/stores/new",
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}
function editStore(store_id) {
    $.ajax({
        async: false,
        type: "get",
        url: "/admin/stores/" + store_id,
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}

function deleteStore(store_id) {
    $.ajax({
        async: false,
        type: "delete",
        url: "/admin/stores/" + store_id,
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}

function openDeleteModal(store_id) {
    $('#mini_modal .header').html("Alert");
    $('#mini_modal .content').html("Are you sure to delete store?");
    $('#mini_modal')
    .modal({
        onApprove : function() {
            deleteStore(store_id)
        }
    })
    .modal('show');
}


</script>
