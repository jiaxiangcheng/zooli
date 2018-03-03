{{template "common/modal.tpl" .}}

<div class="row">
    <div class="column">

        <div class="ui segments">
            <div class="ui segment">
                <h1 class="ui header center aligned">{{i18n .Lang "orders_table.title"}}</h1>
            </div>
            <div class="ui segment">
                {{template "common/flash.tpl" .}}
                <table class="ui compact selectable striped celled table tablet stackable" id="data_table" cellspacing="0" width="100%">
                    <thead>
                        <tr>
                            <th>Serial number</th>
                            <th>{{i18n .Lang "table_attribute_names.customer"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.product"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.status"}}</th>
                            <th>{{i18n .Lang "table_attribute_names.fee"}}</th>
                            <th class="center aligned">Created</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {{ range .orders }}
                        <tr>
                            <td>{{ .ID}}</td>
                            <td>{{ .Customer.Name}}</td>
                            <td>{{ .Product.Name}}</td>
                            <td><a class="ui {{if eq .Status $.ordered}}orange{{else if eq .Status $.inService}}yellow{{else if eq .Status $.endService}}olive{{else if eq .Status $.waitingForPayment}}teal{{else if eq .Status $.orderFinished}}green{{else}}grey{{end}} label">{{.Status}}</a></td>
                            <td>{{ .Fee}}</td>
                            <td class="center aligned">{{ .CreatedAt.Format "2006-01-02 15:04"}}</td>
                            <td class="center aligned">
                                <i class="blue link pencil alternate icon" onclick="editOrder('{{ .ID}}');"></i>
                                <i class="red link trash alternate icon" onclick="openDeleteModal('{{ .ID}}');"></i>
                            </td>
                        </tr>
                        {{ end }}
                    </tbody>
                    <tfoot class="full-width">
                        <tr>
                            <th colspan="5">
                                <div class="ui right floated small primary labeled icon button" onclick="newOrder();">
                                    <i class="payment icon"></i> {{i18n .Lang "orders_table.add_order"}}
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
        order: [[ 5, "desc" ]],
        info: false
    }
);

});

function newOrder() {
    $.ajax({
        async: false,
        type: "get",
        url: "/public/orders/new",
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}

function editOrder(order_id) {
    $.ajax({
        async: false,
        type: "get",
        url: "/public/orders/" + order_id,
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}

function deleteOrder(order_id) {
    $.ajax({
        async: false,
        type: "delete",
        url: "/public/orders/" + order_id,
        success: function (data) {
            $('#main_content').html(data);
        }
    });
}

function openDeleteModal(order_id) {
    $('#mini_modal .header').html("Alert");
    $('#mini_modal .content').html("Are you sure to delete order?");
    $('#mini_modal')
    .modal({
        onApprove : function() {
            deleteOrder(order_id)
        }
    })
    .modal('show');
}

</script>
