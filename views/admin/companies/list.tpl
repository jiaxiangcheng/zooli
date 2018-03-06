{{template "common/modal.tpl" .}}

<div class="row">
    <div class="column">

        <div class="ui segments">
            <div class="ui segment">
                <h1 class="ui header center aligned">{{i18n .Lang "companies_table.title"}}</h1>
            </div>
            <div class="ui segment">
                {{template "common/flash.tpl" .}}
                <table class="ui compact selectable striped celled table tablet stackable" id="data_table" cellspacing="0" width="100%">

                    <thead>
                        <tr>
                            <th>{{i18n .Lang "table_attribute_names.company_name"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.contact"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.phone"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.email"}}</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ range .companies }}
                        <tr>
                            <td>{{ .Name}}</td>
                            <td>{{ .Contact}}</td>
                            <td>{{ .PhoneNumber}}</td>
                            <td>{{ .Email}}</td>
                            <td class="center aligned">
                                <i class="blue link pencil alternate icon" onclick="editCompany('{{ .ID}}');"></i>
                                <i class="red link trash alternate icon" onclick="openDeleteModal('{{ .ID}}');"></i>
                            </td>

                        </tr>
                        {{ end }}
                    </tbody>
                    <tfoot class="full-width">
                        <tr>
                            <th colspan="5">
                                <div class="ui right floated small primary labeled icon button" onclick="newCompany();">
                                    <i class="world icon"></i> {{i18n .Lang "companies_table.add_company"}}
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

function newCompany() {
    $.ajax({
        async: false,
        type: "get",
        url: "/admin/companies/new",
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}

function editCompany(company_id) {
    $.ajax({
        async: false,
        type: "get",
        url: "/admin/companies/" + company_id,
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}

function deleteCompany(company_id) {
    $.ajax({
        async: false,
        type: "delete",
        url: "/admin/companies/" + company_id,
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}

function openDeleteModal(company_id) {
    $('#mini_modal .header').html("Alert");
    $('#mini_modal .content').html("Are you sure to delete company?");
    $('#mini_modal')
    .modal({
        onApprove : function() {
            deleteCompany(company_id)
        }
    })
    .modal('show');
}
</script>
