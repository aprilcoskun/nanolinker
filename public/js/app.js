"use strict";

function openEditRow(id) {
    var idId = "#table_id_" + id;
    var urlId = "#table_url_" + id;

    var oldId = $(idId).text();
    var oldUrl = $(urlId).text();
    $(idId).data("old", oldId);
    $(idId).html(toInput(idId, oldId));
    $(urlId).html(toInput(urlId, oldUrl));
    $("#table_actions_" + id).html(cancelOrSaveButtons(id))
}

function cancelEditRow(id) {
    var idVal = $("#table_id_" + id + "_input").val();
    var urlVal = $("#table_url_" + id + "_input").val();

    $("#table_id_" + id).html("<a href='/l/" + idVal + "'>" + idVal + "</a>");
    $("#table_url_" + id).html("<a href='" + urlVal + "'>" + urlVal + "</a>");
    $("#table_actions_" + id).html(editOrDeleteButtons(id));
}

function saveEditLink(id) {
    var idVal = $("#table_id_" + id + "_input").val();
    var urlVal = $("#table_url_" + id + "_input").val();
    var oldId = $("#table_id_" + id).data("old");

    if (!urlVal || !idVal) {
        return alert("Empty Value(s)");
    }

    if (!oldId) {
        return alert("Old Id not found!!!!")
    }

    $.ajax({
        url: "/v1/link/" + oldId,
        data: JSON.stringify({id: idVal, url: urlVal}),
        contentType: "application/json",
        type: "PUT",
        success: function success() {
            location.href = "/v1";
        },
        error: function error(data, status, err) {
            alert(err);
        }
    });
}

function saveLink() {
    var newId = $("#new_id").val();
    var newUrl = $("#new_url").val();
    if (!newUrl) {
        return alert("Empty Link");
    }
    $.ajax({
        url: "/v1/link",
        data: JSON.stringify({id: newId, url: newUrl}),
        contentType: "application/json",
        type: "POST",
        success: function success() {
            location.href = "/v1";
        },
        error: function error(data, status, err) {
            return alert(err);
        }
    });
}

function openDeleteModal(id) {
    $('.basic.modal').modal("setting", "closeable", false).modal("show");
    $('#delete_button').data("id", id);
}

function deleteRow() {
    var id = $('#delete_button').data("id");
    $.ajax({
        url: "/v1/link/" + id,
        type: "DELETE",
        success: function success() {
            location.href = "/v1";
        },
        error: function error(data, status, err) {
            alert(err);
        }
    })
}

function toLink(val) {
    return "<a href='" + val + "'>" + val + "</a>";
}

function toInput(id, val) {
    return '<div class="ui input small fluid"><input id="' + id.substr(1) + '_input" type="text" value="' + val + '"/></div>';
}

function cancelOrSaveButtons(id) {
    return '<div class="ui buttons">' +
        '<button class="ui green inverted button"  onclick="saveEditLink(\'' + id + '\')"><i class="check icon"></i>Save</button>' +
        '<button class="ui red inverted button" onclick="cancelEditRow(\'' + id + '\')"><i class="close icon"></i>Cancel</button>' +
        '</div>';
}

function editOrDeleteButtons(id) {
    return '<div class="ui buttons">' +
        '<button class="ui icon olive button" onclick="openEditRow(\'' + id + '\')"><i class="pen icon"></i>Edit</button>' +
        '<button class="ui icon red button" onclick="openDeleteModal(\'' + id + '\')"><i class="trash alternate icon"></i>Delete</button>' +
        '</div>';
}

$(function () {
    $('.created-at').each(function (index, element) {
        var textDate = $(this).text();
        $(this).text(new Date(textDate).toLocaleString())
    });
});