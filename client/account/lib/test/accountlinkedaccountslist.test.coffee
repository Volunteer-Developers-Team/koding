kd                        = require 'kd'
expect                    = require 'expect'
KodingListView            = require 'app/kodinglist/kodinglistview'
AccountLinkedAccountsList = require 'account/accountlinkedaccountslist'


describe 'AccountLinkedAccountsList', ->

  describe 'constructor', ->

    it 'should be extended from KodingListView', ->

      listView      = new AccountLinkedAccountsList
      instanceCheck = listView instanceof KodingListView

      expect(instanceCheck).toBeTruthy()

    it 'should be a <ul> HTML element', ->

      listView  = new AccountLinkedAccountsList

      expect(listView.getOptions().tagName).toEqual 'ul'

    it 'should use custom css class', ->

      listView  = new AccountLinkedAccountsList
      hasClass  = listView.getDomElement().hasClass 'AppModal--account-switchList'

      expect(hasClass).toBeTruthy()
