<h1 class="ui header">Stores</h1>
{{template "common/flash.tpl" .}}
<table class="ui single line striped collapsing table">
    <thead>
    <tr>
        <th class="center aligned">Name</th>
        <th class="center aligned">Address</th>
        <th class="center aligned">Latitude</th>
        <th class="center aligned">Longitude</th>
        <th class="center aligned">Phone number</th>
        <th class="center aligned">Company name</th>
        <th class="center aligned">Manager</th>
        <th class="center aligned">Services</th>
        <th class="center aligned"></th>
        <th></th>
    </tr>
    </thead>
    <tbody>
    {{ range .stores }}
    <tr>
        <td class="center aligned">{{ .Name}}</td>
        <td class="center aligned">{{ .Address}}</td>
        <td class="center aligned">{{ .Latitude}}</td>
        <td class="center aligned">{{ .Longitude}}</td>
        <td class="center aligned">{{ .PhoneNumber}}</td>
        <td class="center aligned">{{ .Company.Name}}</td>
        <td class="center aligned">
            <div class="ui {{if .Manager.Name}}primary{{else}}negative{{end}} basic animated fade button" tabindex="0">
                <div class="visible content">{{.Manager.Name}}</div>
                <div class="hidden content">
            {{if .Manager.Name}}Change{{else}}Assign{{end}}
                </div>
            </div>
        </td>
        <td class="center aligned">{{range .Services}} <a class="ui blue label">{{.Name}}</a> {{end}}</td>
        <td class="center aligned">
            <button type="button"
                    class="ui basic button"
                    onclick="editStore('{{ .ID}}');">
                View
            </button>
        </td>

        <td class="center aligned">
            <button type="button"
                    class="ui negative button"
                    data-toggle="modal"
                    data-target=".bs-example-modal-sm"
                    onclick="deleteStore('{{ .ID}}');">
                Delete
            </button>
        </td>

    </tr>
    {{ end }}
    </tbody>
</table>
<button type="button"
        title="View store"
        class="ui basic big button"
        onclick="newStore();"
        style="margin: 10px 10px">
    <i class="add store icon"></i>
    Create store
</button>



<script type="text/javascript">
    function newStore() {
        $.ajax({
            async: false,
            type: "get",
            url: "/stores/new",
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
    function editStore(store_id) {
        $.ajax({
            async: false,
            type: "get",
            url: "/stores/" + store_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
    function deleteStore(store_id) {
        $.ajax({
            async: false,
            type: "delete",
            url: "/stores/" + store_id,
            success: function (data) {
                $('#main_content').html(data);
            }
        });
    }
</script>
