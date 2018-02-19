<h1 class="ui header" style="text-align:center;">Companies</h1>
{{template "common/flash.tpl" .}}
<table class="ui single line striped collapsing table" id="table_list">
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
        <td class="center aligned">{{ .Name}}</td>
        <td class="center aligned">{{ .Contact}}</td>
        <td class="center aligned">{{ .PhoneNumber}}</td>
        <td class="center aligned">{{ .Email}}</td>
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
                    data-toggle="modal"
                    data-target=".bs-example-modal-sm"
                    onclick="deleteCompany('{{ .ID}}');">
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
            title="View company"
            class="ui basic big button"
            onclick="newCompany();">
        <i class="add company icon"></i>
        Create company
    </button>
</div>
<style>
    #table_list {
       margin-left:auto;
       margin-right:auto;
     }
</style>

<script type="text/javascript">
    function newCompany() {
        $.ajax({
            async: false,
            type: "get",
            url: "/companies/new",
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }

    function editCompany(company_id) {
        $.ajax({
            async: false,
            type: "get",
            url: "/companies/" + company_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
    function deleteCompany(company_id) {
        $.ajax({
            async: false,
            type: "delete",
            url: "/companies/" + company_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
</script>
