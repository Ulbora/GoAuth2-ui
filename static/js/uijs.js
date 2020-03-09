
// $(document)
// .ready(function() {

$(function () {
    $('.ui.menu .ui.dropdown').dropdown({
        on: 'hover'
    });
    $('.ui.menu a.item')
        .on('click', function () {
            $(this)
                .addClass('active')
                .siblings()
                .removeClass('active')
                ;
        })
        ;
});
// })
// ;

$(function () {
    $('a.item').click(function () {
        $('.item').removeClass('active');
        $(this).addClass('active');
    })
});


$(function () {
    $('.activating.element')
        .popup()
        ;
});



$(function () {
    $('.ui.radio.checkbox')
        .checkbox()
        ;
});


$(function () {
    $('.ui.checkbox')
        .checkbox()
        ;
});

var urlId;
var urlName;
var showModel = function (id, name) {
    urlId = id;
    urlName = name;
    $(function () {
        $('.ui.modal').modal('show');
    });
    $(".ui.modal").modal({
        closable: true
    });
    document.getElementById("cont").innerHTML = "Delete URL " + name + "?";
}

var getUrlName = function () {
    return urlName;
}
var getRedirectUriId = function () {
    return urlId;
}


// $(function () {
//     $('.ui.modal')
//     .modal()
// });

$(function () {
    $("#delete").click(function () {
        $(".ui.modal").modal('show');
    });
    $(".ui.modal").modal({
        closable: true
    });
});



var grantTypeId;
var grantType;
var showGrantTypeModel = function (id, gt) {
    grantTypeId = id;
    grantType = gt;
    $(function () {
        $('.ui.modal').modal('show');
    });
    $(".ui.modal").modal({
        closable: true
    });
    document.getElementById("cont").innerHTML = "Delete Grant Type " + gt + "?";
}

var getGrantTypeId = function () {
    return grantTypeId;
}



var roleId;
var clientRole;
var showRoleModel = function (id, rl) {
    roleId = id;
    clientRole = rl;
    $(function () {
        $('.ui.modal').modal('show');
    });
    $(".ui.modal").modal({
        closable: true
    });
    document.getElementById("cont").innerHTML = "Delete Client Role " + rl + "?";
}

var getRoleId = function () {
    return roleId;
}


var uriId;
var clientUir;
var allowedUriRoleId
var showAllowedUriModel = function (id, ri, rid) {
    uriId = id;
    clientUir = ri;
    allowedUriRoleId = rid;
    $(function () {
        $('.ui.modal').modal('show');
    });
    $(".ui.modal").modal({
        closable: true
    });
    document.getElementById("cont").innerHTML = "Delete Client URI " + ri + "?";
}

var getAllowedUriId = function () {
    return uriId;
}

var getAllowedUriRoleId = function () {
    return allowedUriRoleId;
}


var unHideAllowedUri = function (id) {
    document.getElementById(id).disabled = false
    document.getElementById("del" + id).style.display = 'none';
    document.getElementById("sub" + id).style.display = 'block';
}

var ulborUriAddBntDisable = function () {
    document.getElementById("ulborUriAddBnt").disabled = false
    document.getElementById("progBar").style.display = 'block';
    $('#progBar')
        .progress('increment')
        ;
    setInterval(function () {
        $('#progBar')
            .progress('increment')
            ;
    }, 500)
}


// $(function(){
// 	$('#progBar')
//     .progress('increment')    
//   ;
// });
var addUserFormVisible = false;
var showAddUserForm = function () {
    if (addUserFormVisible === false) {
        document.getElementById("addUserForm").style.display = 'block';
        addUserFormVisible = true;
    } else {
        document.getElementById("addUserForm").style.display = 'none';
        addUserFormVisible = false;
    }

}


var editUserInfoFormVisible = false;
var showEditUserInfoForm = function () {
    if (editUserInfoFormVisible === false) {
        document.getElementById("editUserInfoForm").style.display = 'block';
        editUserInfoFormVisible = true;
    } else {
        document.getElementById("editUserInfoForm").style.display = 'none';
        editUserInfoFormVisible = false;
    }

}


var editUserEnableFormVisible = false;
var showEditUserEnableForm = function () {
    if (editUserEnableFormVisible === false) {
        document.getElementById("editUserEnableForm").style.display = 'block';
        editUserEnableFormVisible = true;
    } else {
        document.getElementById("editUserEnableForm").style.display = 'none';
        editUserEnableFormVisible = false;
    }

}


var editUserPwFormVisible = false;
var showEditUserPwForm = function () {
    if (editUserPwFormVisible === false) {
        document.getElementById("editUserPwForm").style.display = 'block';
        editUserPwFormVisible = true;
    } else {
        document.getElementById("editUserPwForm").style.display = 'none';
        editUserPwFormVisible = false;
    }

}


var showAddGatewayAccountForm = function () {
    document.getElementById("addGatewayAccountMsg").style.display = 'none';
    document.getElementById("addGatewayAccountForm").style.display = 'block';
}


var gwRouteId;
var gwRoute;
var showGwRouteModel = function (id, name) {
    gwRouteId = id;
    gwRoute = name;
    $(function () {
        $('.ui.modal').modal('show');
    });
    $(".ui.modal").modal({
        closable: true
    });
    document.getElementById("cont").innerHTML = "Delete Gateway Route " + name + "?";
}


var getGwRoute = function () {
    return gwRoute;
}
var getGwRouteId = function () {
    return gwRouteId;
}


var gwRouteUrlId;
var urlName;
var rtUrl
var showGwRouteUrlModel = function (id, name, url) {
    document.getElementById("del").style.display = 'block';
    document.getElementById("act").style.display = 'none';
    gwRouteUrlId = id;
    urlName = name;
    rtUrl = url;
    $(function () {
        $('.ui.modal').modal('show');
    });
    $(".ui.modal").modal({
        closable: true
    });
    document.getElementById("hd").innerHTML = "Delete this URL?";
    document.getElementById("cont").innerHTML = "Delete Route URL Name: " + name + "; URL: " + rtUrl+ " ?";
    
}


var getGwRouteUrlId = function () {
    return gwRouteUrlId;
}



var gwRouteUrlId2;
//var urlName2;
//var rtUrl2
var showGwRouteUrlActiveModel = function (id, name, url) {
    document.getElementById("act").style.display = 'block';
    document.getElementById("del").style.display = 'none';
    gwRouteUrlId2 = id;
    //urlName2 = name;
   // rtUrl2 = url;
    $(function () {
        $('.ui.modal').modal('show');
    });
    $(".ui.modal").modal({
        closable: true
    });
    document.getElementById("hd").innerHTML = "Activate URL?";
    document.getElementById("cont").innerHTML = "Activate URL: " + name + " " + url+ " ?";
    
}


var getGwRouteUrlId2 = function () {
    return gwRouteUrlId2;
}

var showCb = false;
var showCircuitBreaker = function (cbEnabled) {
    if(showCb == false){
        showCb = cbEnabled;
    }    
    if(showCb === false){
        document.getElementById("cb").style.display = 'block';
        //document.getElementById("onoff").innerHTML = "On";
        showCb = true;
    }else{
        document.getElementById("cb").style.display = 'none';
       // document.getElementById("onoff").innerHTML = "Off";
        showCb = false;
    }    
}