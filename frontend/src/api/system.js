export function getCMDList() {
    return [
        {
            cmd: "calc", //命令
            param: '',//命令参数
            name: "计算器",
            icon: require('../assets/system/calc.png'),
            desc: "计算器"
        },
        {
            cmd: "explorer", //命令
            param: 'D:\\', //指定目录
            name: "打开目录",
            icon: require('../assets/system/explorer.png'),
            desc: "以资源管理器方式，打开目录"
        },
        {
            cmd: "control", //命令
            param: '',
            name: "控制面板",
            icon: require('../assets/system/control.png'),
            desc: "控制面板"
        }, {
            cmd: "osk", //命令
            param: '',
            name: "屏幕键盘",
            icon: require('../assets/system/osk.png'),
            desc: "屏幕键盘"
        },
        {
            cmd: "notepad", //命令
            param: '',
            name: "记事本",
            icon: require('../assets/system/notepad.png'),
            desc: "记事本"
        },
        {
            cmd: "devmgmt.msc", //命令
            param: '',
            name: "设备管理",
            icon: require('../assets/system/devmgmt.png'),
            desc: "设备管理"
        },
        {
            cmd: "dxdiag", //命令
            param: '',
            name: "DirectX",
            icon: require('../assets/system/dxdiag.png'),
            desc: "检查DirectX信息"
        },
        {
            cmd: "Msconfig", //命令
            param: '',
            name: "系统配置",
            icon: require('../assets/system/msconfig.png'),
            desc: "系统配置"
        },
        {
            cmd: "regedit", //命令
            param: '',
            name: "注册表",
            icon: require('../assets/system/regedit.png'),
            desc: "注册表"
        },
        {
            cmd: "taskmgr", //命令
            param: '',
            name: "任务管理",
            icon: require('../assets/system/taskmgr.png'),
            desc: "任务管理"
        },
        {
            cmd: "mspaint", //命令
            param: '',
            name: "画图",
            icon: require('../assets/system/mspaint.png'),
            desc: "画图"
        },
        {
            cmd: "mstsc", //命令
            param: '',
            name: "远程桌面",
            icon: require('../assets/system/mstsc.png'),
            desc: "远程桌面"
        },
        {
            cmd: "magnify", //命令
            param: '',
            name: "放大镜",
            icon: require('../assets/system/magnify.png'),
            desc: "放大镜"
        },
        {
            cmd: "diskmgmt.msc", //命令
            param: '',
            name: "磁盘管理",
            icon: require('../assets/system/diskmgmt.png'),
            desc: "磁盘管理"
        }
    ]
}
