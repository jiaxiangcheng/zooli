<h1 class="ui header" style="text-align:center;">Stores</h1>
{{template "common/modal.tpl" .}}
{{template "common/flash.tpl" .}}
<table class="ui single line striped collapsing table" id="table_list" style="table-layout:fixed; width:100%;">
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
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Name}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Address}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Latitude}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Longitude}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .PhoneNumber}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{ .Company.Name}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{range .Managers}} <a class="ui olive label">{{.Name}}</a>{{end}}</td>
        <td class="center aligned" style="overflow: hidden;text-overflow: ellipsis;">{{range .Services}} <a class="ui blue label">{{.Name}}</a> {{end}}</td>
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

<div class="ui middle aligned center aligned grid">
    <button type="button"
            title="View store"
            class="ui basic big button"
            onclick="newStore();"
            style="margin: 15px;">
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
