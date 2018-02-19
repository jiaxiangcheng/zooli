<h1 class="ui header" style="text-align:center;">Stores</h1>
{{template "common/modal.tpl" .}}
{{template "common/flash.tpl" .}}
<table class="ui single line striped collapsing table" id="table_list">
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
    {{ range $i, $s := .stores }}
    <tr>
        <td class="center aligned">{{ .Name}}</td>
        <td class="center aligned">{{ .Address}}</td>
        <td class="center aligned">{{ .Latitude}}</td>
        <td class="center aligned">{{ .Longitude}}</td>
        <td class="center aligned">{{ .PhoneNumber}}</td>
        <td class="center aligned">{{ .Company.Name}}</td>
        <td class="center aligned">
            <div class="ui {{if .Manager.Name}}green{{else}}yellow{{end}} floating dropdown icon button">
                {{if .Manager.Name}} {{else}} <i class="add user icon"></i> {{end}}
                <span class="text">{{.Manager.Name}}</span>
                <div class="menu">
                    <div class="ui icon search input">
                        <i class="search icon"></i>
                        <input type="text" placeholder="Search managers...">
                    </div>
                    <div class="scrolling menu">
                        {{range $.managers}}
                        <div class="item" data-value="{{.ID}},{{$s.ID}}">
                            <span class="text">{{.Username}}</span>
                            <span class="text">({{.Name}})</span>
                        </div>
                         {{end}}
                    </div>
                </div>
            </div>
            <!--div class="ui {{if .Manager.Name}}primary{{else}}negative{{end}} basic animated fade button" tabindex="0" onclick="openAssignmentModal({{.ID}},{{.Manager.ID}})">
                <div class="visible content">{{.Manager.Name}}</div>
                <div class="hidden content">
            {{if .Manager.Name}}Change{{else}}Assign{{end}}
                </div>
            </div-->
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
                    onclick="openDeleteModal('{{ .ID}}');">
                Delete
            </button>
        </td>

    </tr>
    {{ end }}
    </tbody>
</table>
<!--div class="ui tiny modal" id="user_modal">
    <i class="close icon"></i>
    <div class="header">Pick a manager</div>
    <div class="scrolling content">
        <div style="min-height:400px;" >
            <div class="ui grid container">
                {{range .managers}}
                    <div class="four wide column"><span>{{.Name}}</span></div>
                {{end}}
            </div>
        </div>
    </div>
    <div class="actions">
        <div class="ui deny button">
            Cancel
        </div>
        <div class="ui positive button">
            Save
        </div>
    </div>
</div-->

</br>
<div class="ui middle aligned center aligned grid">
    <button type="button"
            title="View store"
            class="ui basic big button"
            onclick="newStore();">
        <i class="add store icon"></i>
        Create store
    </button>
</div>

<style>
    #table_list {
       margin-left:auto;
       margin-right:auto;
     }
</style>

<script type="text/javascript">
    $(document)
            .ready(function() {
                $('.dropdown')
                    .dropdown({
                        action: function(text, value) {
                            args = value.split(",");
                            $.ajax({
                                type: "post",
                                url: "/users/" + args[0] + "/assign",
                                data: {"storeID": args[1]},
                                success: function (data) {
                                    $('#main_content').html(data);
                                }
                            });
                        }
                    });
            });


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

    function openAssignmentModal(storeID, userID) {
        $('#user_modal')
            .modal({
                onVisible: function () {

                }
            })
            .modal("show");
    }
</script>
