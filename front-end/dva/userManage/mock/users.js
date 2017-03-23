'use strict';

const qs = require('qs');

const mockjs = require('mockjs');

module.exports = {
  'GET /api/users' (req, res) {
    const page = qs.parse(req.query);

    const data = mockjs.mock({
      'data|100': [{
        'id|+1': 1,
        name: '@cname',
        address: '@region'
      }],
      page: {
        total: 10,
        current: 1
      }
    });

    res.json({
      success: true,
      data: data.data,
      page: data.page
    })
  }
}
