module.exports = [
  {
    key: 'users',
    name: '用户',
    icon: 'user',
  },
  {
    key: 'sys',
    name: '系统设置',
    icon: 'setting',
    child: [

      {
        key: 'record_libraries',
        path: 'record_libraries',
        name: '解析库',
        icon: 'cloud',
        visual: true,
      },
      {
        key: 'user_groups',
        path: 'user_groups',
        name: '用户组',
        icon: 'team'
      },
      {
        key: 'layer_tree',
        path: 'layer_tree',
        name: '接入层定义',
        icon: 'share-alt'
      },
      {
        key: 'layers',
        path: 'layers',
        name: '接入层定义',
        visual: true,
      },
      {
        key: 'domains',
        path: 'domains',
        name: '根域名列表',
        visual: true
      },
      {
        key: 'departments',
        path: 'departments',
        name: '部门定义',
        visual: true
      },
      {
        key: 'record_lines',
        path: 'record_lines',
        name: '线路类型',
        visual: true
      },
      {
        key: 'record_types',
        path: 'record_types',
        name: '记录类型',
        visual: true
      }
    ]
  },
];
