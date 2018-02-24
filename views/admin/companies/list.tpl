<h1 class="ui header" style="text-align:center;">Companies</h1>
{{template "common/modal.tpl" .}}
{{template "common/flash.tpl" .}}
<div class="ui divider"></div>

<table class="ui single line striped collapsing table" id="companies-table"
    style="margin-left:auto; margin-right:auto; table-layout:fixed; width:100%;">
    <thead>
    <tr>
        <th class="center aligned">Name</th>
        <th class="center aligned">Contact</th>
        <th class="center aligned">Phone number</th>
        <th class="center aligned">Email</th>
        <th class="center aligned"></th>
        <th></th>
    </tr>
    </thead>
    <tbody>
    {{ range .companies }}
    <tr>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Name}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Contact}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .PhoneNumber}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Email}}</td>
        <td class="center aligned">
            <button type="button"
                    class="ui basic button"
                    onclick="editCompany('{{ .ID}}');">
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
            title="View company"
            class="ui blue basic big button"
            onclick="newCompany();"
            style="margin: 15px;">
        <i class="add icon"></i>
        Create company
    </button>
</div>


<script type="text/javascript">


    $(document).ready(function() {
        $('#companies-table').DataTable();
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
