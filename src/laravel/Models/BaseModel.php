<?php

namespace CAN\Models;

use MongoDB\Laravel\Eloquent\Model;
use MongoDB\Laravel\Eloquent\SoftDeletes;
use Wildside\Userstamps\Userstamps;

abstract class BaseModel extends Model
{
    use Userstamps;
}
