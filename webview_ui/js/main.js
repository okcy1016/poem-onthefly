function hideEle(id) {
    var x = document.getElementById(id);
    x.style.display = "none";
}

function showEle(id) {
    var x = document.getElementById(id);
    x.style.display = "block";
}

function showCorr(key) {
    if (key === "1") {
	showEle("intro");
    } else {
	hideEle("intro");
    }
    
    if (key === "2-1") {
	showEle("fontSlider");
    }
    else {
	hideEle("fontSlider");
    }
    
    if (key === "2-3") {
	showEle("intervalSlider");
    }
    else {
	hideEle("intervalSlider");
    }

    if (key === "2-2") {
	showEle("imgSel");
    } else {
	hideEle("imgSel");
    }

    if (key === "2-4") {
	showEle("chkBox0");
    }
    else {
	hideEle("chkBox0");
    }
}

function initEle() {
    hideEle("fontSlider");
    hideEle("intervalSlider");
    hideEle("imgSel");
    hideEle("chkBox0")
    showEle("intro");
}

var cfgBak = {
    "color": "(0,0,0,0.8)",
    "font": "fonts/FangZhengKaiTiJianTi-1.ttf",
    "fontSize": "25",
    "imgLoc": "webview_ui/images/11.jpg",
    "interval": "1",
    "xOffset": "+0",
    "yOffset": "+0",
    "ifDae": false
};

var Main = {
    data: function() {
	return {
            activeIndex: "1",
            // Get config from config.js
            fontSize: Number(cfg.fontSize),
            fontColor: "rgba" + cfg.color,
            predefineColors: [
		"#ff4500",
		"#ff8c00",
		"#ffd700",
		"#90ee90",
		"#00ced1",
		"#1e90ff",
		"#c71585",
		"rgba(255, 69, 0, 0.68)",
		"rgb(255, 120, 0)",
		"hsv(51, 100, 98)",
		"hsva(120, 40, 94, 0.5)",
		"hsl(181, 100%, 37%)",
		"hsla(209, 100%, 56%, 0.73)",
		"#c7158577"
            ],
            interval: Number(cfg.interval),
            imgSelOptions: [{
		value: "0.jpg"
            }, {
		value: "1.jpg"
	    }, {
		value: "2.jpg"
	    }, {
		value: "3.jpg"
	    }, {
		value: "4.jpg"
	    }, {
		value: "5.jpg"
	    }, {
		value: "6.jpg"
	    }, {
		value: "7.jpg"
	    }, {
		value: "8.jpg"
	    }, {
		value: "9.jpg"
	    }, {
		value: "10.jpg"
	    }, {
		value: "11.jpg"
	    }, {
		value: "12.jpg"
	    }, {
		value: "13.jpg"
            }, {
		value: "14.jpg"
            }, {
		value: "15.jpg"
            }],
            // Get imgloc from config.js
            imgSelVal: cfg.imgLoc.substring(18, cfg.imgLoc.length),

	    // Initial status of checkbox 0
	    chkBox0Chked: cfg.ifDae

	}
    },
    
    methods: {
	
	handleSelect: function(key) {
            showCorr(key);
	},
	
	fontSizeChanged: function(fontSize) {
            cfg.fontSize = fontSize.toString();
	},
	
	fontColorChanged: function(color) {

            // Remove the "rgba" string
            cfg.color = color.substring(4, color.length);
	},
	
	intervalChanged: function(interval) {
            cfg.interval = interval.toString();
	},

	imgSelChgd: function(value) {
            cfg.imgLoc = "webview_ui/images/" + value;
	},

	saveAndRestartService: function() {

            // Remove spaces in color string
            // Invoke external Go function
            external.invoke(JSON.stringify(cfg).replace(/\s+/g, ""));
	},

	chkBox0Chgd: function(ifChked) {
	    cfg.ifDae = ifChked;
	}
    }
}

initEle();
var Ctor = Vue.extend(Main)
new Ctor().$mount("#app")
