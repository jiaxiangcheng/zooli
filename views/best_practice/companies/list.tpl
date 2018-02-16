
<h1 class="ui header">Companies</h1>
{{template "best_practice/common/flash.tpl" .}}
<table class="ui single line striped collapsing table">
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
        <!-->{{if ne $.user.ID .ID}}
        <td class="center aligned">
            <button type="button"
                    class="ui negative button"
                    data-toggle="modal"
                    data-target=".bs-example-modal-sm"
                    onclick="deleteCompany('{{ .ID}}');">
                Delete
            </button>
        </td>
        {{end}}<-->
    </tr>
    {{ end }}
    </tbody>
</table>
<button type="button"
        title="View company"
        class="ui basic big button"
        onclick="newCompany();"
        style="margin: 10px 10px">
    <i class="add company icon"></i>
    Create company
</button>

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
    }/*
    function deleteUser(user_id) {
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
