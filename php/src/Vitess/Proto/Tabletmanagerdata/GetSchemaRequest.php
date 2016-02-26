<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: tabletmanagerdata.proto
//   Date: 2016-01-22 01:34:35

namespace Vitess\Proto\Tabletmanagerdata {

  class GetSchemaRequest extends \DrSlump\Protobuf\Message {

    /**  @var string[]  */
    public $tables = array();
    
    /**  @var boolean */
    public $include_views = null;
    
    /**  @var string[]  */
    public $exclude_tables = array();
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'tabletmanagerdata.GetSchemaRequest');

      // REPEATED STRING tables = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "tables";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_REPEATED;
      $descriptor->addField($f);

      // OPTIONAL BOOL include_views = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "include_views";
      $f->type      = \DrSlump\Protobuf::TYPE_BOOL;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // REPEATED STRING exclude_tables = 3
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 3;
      $f->name      = "exclude_tables";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_REPEATED;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <tables> has a value
     *
     * @return boolean
     */
    public function hasTables(){
      return $this->_has(1);
    }
    
    /**
     * Clear <tables> value
     *
     * @return \Vitess\Proto\Tabletmanagerdata\GetSchemaRequest
     */
    public function clearTables(){
      return $this->_clear(1);
    }
    
    /**
     * Get <tables> value
     *
     * @param int $idx
     * @return string
     */
    public function getTables($idx = NULL){
      return $this->_get(1, $idx);
    }
    
    /**
     * Set <tables> value
     *
     * @param string $value
     * @return \Vitess\Proto\Tabletmanagerdata\GetSchemaRequest
     */
    public function setTables( $value, $idx = NULL){
      return $this->_set(1, $value, $idx);
    }
    
    /**
     * Get all elements of <tables>
     *
     * @return string[]
     */
    public function getTablesList(){
     return $this->_get(1);
    }
    
    /**
     * Add a new element to <tables>
     *
     * @param string $value
     * @return \Vitess\Proto\Tabletmanagerdata\GetSchemaRequest
     */
    public function addTables( $value){
     return $this->_add(1, $value);
    }
    
    /**
     * Check if <include_views> has a value
     *
     * @return boolean
     */
    public function hasIncludeViews(){
      return $this->_has(2);
    }
    
    /**
     * Clear <include_views> value
     *
     * @return \Vitess\Proto\Tabletmanagerdata\GetSchemaRequest
     */
    public function clearIncludeViews(){
      return $this->_clear(2);
    }
    
    /**
     * Get <include_views> value
     *
     * @return boolean
     */
    public function getIncludeViews(){
      return $this->_get(2);
    }
    
    /**
     * Set <include_views> value
     *
     * @param boolean $value
     * @return \Vitess\Proto\Tabletmanagerdata\GetSchemaRequest
     */
    public function setIncludeViews( $value){
      return $this->_set(2, $value);
    }
    
    /**
     * Check if <exclude_tables> has a value
     *
     * @return boolean
     */
    public function hasExcludeTables(){
      return $this->_has(3);
    }
    
    /**
     * Clear <exclude_tables> value
     *
     * @return \Vitess\Proto\Tabletmanagerdata\GetSchemaRequest
     */
    public function clearExcludeTables(){
      return $this->_clear(3);
    }
    
    /**
     * Get <exclude_tables> value
     *
     * @param int $idx
     * @return string
     */
    public function getExcludeTables($idx = NULL){
      return $this->_get(3, $idx);
    }
    
    /**
     * Set <exclude_tables> value
     *
     * @param string $value
     * @return \Vitess\Proto\Tabletmanagerdata\GetSchemaRequest
     */
    public function setExcludeTables( $value, $idx = NULL){
      return $this->_set(3, $value, $idx);
    }
    
    /**
     * Get all elements of <exclude_tables>
     *
     * @return string[]
     */
    public function getExcludeTablesList(){
     return $this->_get(3);
    }
    
    /**
     * Add a new element to <exclude_tables>
     *
     * @param string $value
     * @return \Vitess\Proto\Tabletmanagerdata\GetSchemaRequest
     */
    public function addExcludeTables( $value){
     return $this->_add(3, $value);
    }
  }
}

